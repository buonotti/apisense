package controllers

import (
	"github.com/buonotti/apisense/api/validation"
	"github.com/gin-gonic/gin"
)

func requestValid(ctx *gin.Context, request any) bool {
	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.AbortWithStatusJSON(400, ErrorResponse{Message: err.Error()})
		return false
	}
	err := validation.FillDefaults(request)
	if err != nil {
		ctx.AbortWithStatusJSON(400, ErrorResponse{Message: err.Error()})
		return false
	}
	if err := validation.Validate(request); err != nil {
		ctx.AbortWithStatusJSON(400, ErrorResponse{Message: err.Error()})
		return false
	}
	return true
}
