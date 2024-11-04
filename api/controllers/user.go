package controllers

import (
	"net/http"

	"github.com/buonotti/apisense/v2/api/db"
	"github.com/gofiber/fiber/v2"
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
func LoginUser(c *fiber.Ctx) error {
	var request LoginRequest
	if err := c.BodyParser(request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(ErrorResponse{Message: err.Error()})
	}

	userData, err := db.LoginUser(request.Username, request.Password)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(ErrorResponse{Message: err.Error()})
	}

	return c.JSON(LoginResponse{Token: userData.Token})
}
