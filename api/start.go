package api

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	cache "github.com/chenyahui/gin-cache"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/buonotti/odh-data-monitor/api/controllers"
	"github.com/buonotti/odh-data-monitor/api/middleware"
	"github.com/buonotti/odh-data-monitor/docs"
	"github.com/buonotti/odh-data-monitor/log"
)

func Start() error {
	// TODO config
	docs.SwaggerInfo.BasePath = "/api"
	gin.SetMode(gin.ReleaseMode)

	store := persist.NewMemoryStore(2 * time.Minute)

	router := gin.New()
	router.Use(middleware.CORS())
	router.Use(log.GinLogger())
	router.Use(gin.Recovery())
	router.Use(middleware.Limiter())
	router.GET("/health", controllers.GetHealth)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	api.GET("/reports", cache.CacheByRequestPath(store, 5*time.Minute), controllers.AllReports)
	api.GET("/reports/:id", cache.CacheByRequestPath(store, 5*time.Minute), controllers.Report)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	done := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		if err := srv.Shutdown(context.Background()); err != nil {
			panic(err) // TODO
		}
		log.ApiLogger.Info("shut down server")
		close(done)
	}()

	log.ApiLogger.Info("server listening")
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}

	<-done
	return nil
}
