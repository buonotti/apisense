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

func Start(host string, port int) error {
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
	api.POST("/login", controllers.LoginUser)
	api.GET("/reports", controllers.AllReports)
	api.GET("/reports/:id", controllers.Report)
	api.GET("/ws", controllers.Ws)

	api.Use(middleware.Auth())
	{
		api.GET("/definitions", controllers.AllDefinitions)
		api.POST("/definitions", controllers.CreateDefinition)
		api.GET("/definitions/:id", controllers.Definition)
	}

	apiHost := host
	if apiHost == "" && viper.GetString("api.host") != "" {
		apiHost = viper.GetString("api.host")
	}

	apiPort := port
	if apiPort == 8080 && viper.GetInt("api.port") != 0 {
		apiPort = viper.GetInt("api.port")
	}

	addr := fmt.Sprintf("%s:%d", apiHost, apiPort)

	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.ApiLogger.WithError(err).Error("cannot start api service")
			} else {
				log.ApiLogger.Info("server stopped")
			}
		}
	}()

	log.ApiLogger.WithField("url", fmt.Sprintf("http://localhost:%v", apiPort)).Info("server listening")

	<-done

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	log.ApiLogger.Info("gracefully stopping server")

	if err := srv.Shutdown(ctx); err != nil {
		err = errors.CannotStopApiServiceError.Wrap(err, "cannot stop server")
	}

	return nil
}
