package middleware

import (
	"net/http"

	"github.com/buonotti/apisense/api/db"
	"github.com/buonotti/apisense/api/jwt"
	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/log"
	"github.com/gofiber/fiber/v2"
	goJWT "github.com/golang-jwt/jwt/v4"
)

// Auth is a middleware that checks if the request has a valid JWT token and then authorizes the request
func Auth() func(c *fiber.Ctx) error {
	const BearerSchema = "Bearer "
	return func(c *fiber.Ctx) error {
		// Get the token string from the header
		authHeader := c.Get("Authorization")
		if authHeader == "" || len(authHeader) <= len(BearerSchema)+1 {
			err := errors.TokenError.New("missing auth token")
			return c.Status(http.StatusUnauthorized).JSON(map[string]any{"message": err.Error()})
		}

		// Cut off the "Bearer " part of the header
		tokenString := authHeader[len(BearerSchema):]

		// Parse and validate the token
		token, err := jwt.ValidateToken(tokenString)
		if err == nil && token.Valid {
			claims := token.Claims.(goJWT.MapClaims)
			uid := claims["uid"].(string)
			log.ApiLogger().Info("Authorizing user", "user", claims["uid"].(string))
			if !db.IsUserEnabled(uid) {
				return c.Status(http.StatusForbidden).JSON(map[string]any{"message": "user is not allowed to use the API"})
			}
		} else {
			return c.Status(http.StatusUnauthorized).JSON(map[string]any{"message": err.Error()})
		}
		return c.Next()
	}
}
