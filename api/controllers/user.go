package controllers

import (
	"net/http"

	"github.com/buonotti/apisense/api/db"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username" validation:"required"`
	Password string `json:"password" validation:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

// LoginUser godoc
//
//	@Summary		Logs a user in
//	@Schemes		LoginRequest LoginResponse ErrorResponse
//	@Description	Logs a user in using the provided credentials
//	@ID				login-user
//	@Tags			authentication
//	@Accept			application/json
//	@Produces		application/json
//	@Param			data	body		LoginRequest	true	"content"
//	@Success		200		{object}	LoginResponse
//	@Failure		400		{object}	ErrorResponse
//	@Failure		401		{object}	ErrorResponse
//	@Failure		500		{object}	ErrorResponse
//	@Router			/login [post]
func LoginUser(c *gin.Context) {
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	userData, err := db.LoginUser(request.Username, request.Password)
	if err != nil {
		c.AbortWithStatusJSON(400, ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(200, LoginResponse{Token: userData.Token})
}
