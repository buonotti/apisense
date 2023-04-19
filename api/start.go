package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/buonotti/apisense/api/controllers"
	"github.com/buonotti/apisense/api/middleware"
	"github.com/buonotti/apisense/docs"
	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/log"
)

func Start() error {
	docs.SwaggerInfo.BasePath = "/api"

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(log.GinMiddleware())
	router.Use(middleware.CORS())
	// router.Use(middleware.Limiter())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	api := router.Group("/api")
	api.GET("/health", controllers.GetHealth)
	api.GET("/reports", controllers.AllReports)
	api.GET("/reports/:id", controllers.Report)
	api.GET("/definitions", controllers.AllDefinitions)
	api.GET("/definitions/:id", controllers.Definition)
	api.GET("/ws", controllers.Ws)

	addr := fmt.Sprintf("%s:%d", viper.GetString("api.host"), viper.GetInt("api.port"))

	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Kill)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.ApiLogger.WithError(err).Error("cannot start api service")
			} else {
				log.ApiLogger.Info("api service stopped")
			}
		}
	}()

	log.ApiLogger.Infof("api service started listening on http://localhost:%v", viper.GetInt("api.port"))

	<-done

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	log.ApiLogger.Info("stopping api service")

	if err := srv.Shutdown(ctx); err != nil {
		err = errors.CannotStopApiServiceError.Wrap(err, "cannot stop api service")
	}

	return nil
}
