package api

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/buonotti/apisense/api/controllers"
	"github.com/buonotti/apisense/api/middleware"
	"github.com/buonotti/apisense/docs"
	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/log"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/spf13/viper"
)

// Start starts the web api accoring to configuation. host and port can be used to override the values in the config
func Start(host string, port int) error {
	if viper.GetString("api.signing_key") == "" {
		return errors.MissingSigningKeyError.New("Missing api.signing_key value in either config or secrets file")
	}
	docs.SwaggerInfo.BasePath = "/api"

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
	})
	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(log.NewFiber())

	app.Get("/swagger/*", swagger.New())
	app.Get("/health", controllers.GetHealth)

	api := app.Group("/api")

	api.Get("/health", controllers.GetHealth)
	api.Post("/login", controllers.LoginUser)
	api.Get("/reports", controllers.AllReports)
	api.Get("/reports/:id", controllers.Report)
	api.Get("/swagger", func(c *fiber.Ctx) error {
		return c.Redirect("/swagger/index.html", http.StatusMovedPermanently)
	})

	if viper.GetBool("api.auth") {
		api.Use(middleware.Auth())
	}
	api.Get("/definitions", controllers.AllDefinitions)
	api.Post("/definitions", controllers.CreateDefinition)
	api.Get("/definitions/:id", controllers.Definition)

	apiHost := host
	if apiHost == "" && viper.GetString("api.host") != "" {
		apiHost = viper.GetString("api.host")
	}

	apiPort := port
	if apiPort == 8080 && viper.GetInt("api.port") != 0 {
		apiPort = viper.GetInt("api.port")
	}

	addr := fmt.Sprintf("%s:%d", apiHost, apiPort)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	go func() {
		_ = <-done
		log.ApiLogger().Info("Interrupt. Stopping server")
		_ = app.Shutdown()
	}()

	log.ApiLogger().Info("Starting server", "address", addr, "handlers", app.HandlersCount())

	err := app.Listen(addr)
	if err != nil {
		return errors.CannotStopApiServiceError.Wrap(err, "failed to start server")
	}

	return nil
}
