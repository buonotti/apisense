package controllers

import (
	"github.com/gofiber/fiber/v2"
)

// GetHealth godoc
//
//	@Summary		Health check
//	@Description	Get the health status of the API
//	@ID				health
//	@Tags			health
//	@Produce		json
//	@Success		200
//	@Router			/health [get]
func GetHealth(c *fiber.Ctx) error {
	return c.JSON(map[string]any{"message": "ok"})
}
