package api

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang.org/x/sys/unix"

	"github.com/buonotti/odh-data-monitor/api/controllers"
	"github.com/buonotti/odh-data-monitor/api/middleware"
	"github.com/buonotti/odh-data-monitor/docs"
	"github.com/buonotti/odh-data-monitor/errors"
	"github.com/buonotti/odh-data-monitor/log"
)

func Start() error {
	// TODO config
	docs.SwaggerInfo.BasePath = "/api"
	gin.SetMode(gin.ReleaseMode)

	// store := persist.NewMemoryStore(2 * time.Minute)

	router := gin.New()
	router.Use(middleware.CORS())
	router.Use(log.GinLogger())
	router.Use(gin.Recovery())
	router.Use(middleware.Limiter())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	api.GET("/health", controllers.GetHealth)
	api.GET("/reports" /*cache.CacheByRequestPath(store, 1*time.Minute),*/, controllers.AllReports)
	api.GET("/reports/:id" /*cache.CacheByRequestPath(store, 1*time.Minute),*/, controllers.Report)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, unix.SIGINT, unix.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			errors.HandleError(err)
		}
	}()

	log.ApiLogger.Info("Api service started")

	<-done

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	log.ApiLogger.Info("Stopping api service")

	if err := srv.Shutdown(ctx); err != nil {
		err = errors.CannotStopApiServiceError.Wrap(err, "Cannot stop api service")
	}
	return nil
}
