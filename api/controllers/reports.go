package controllers

import (
	"github.com/gin-gonic/gin"
)

func AllReports(c *gin.Context) {
	c.JSON(200, gin.H{"message": "ok"})
}

func Report(c *gin.Context) {
	c.JSON(200, gin.H{"message": "ok"})
}
