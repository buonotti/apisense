package controllers

import (
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// GetHealth godoc
// @Summary Health check
// @Description Get the health status of the API
// @ID health
// @Tags health
// @Success 200
// @Router /health [get]
func GetHealth(c *gin.Context) {
	c.JSON(200, gin.H{"message": "ok"})
}
