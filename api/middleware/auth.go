package middleware

import (
	"github.com/buonotti/apisense/api/db"
	"net/http"

	"github.com/buonotti/apisense/api/jwt"
	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/log"
	"github.com/gin-gonic/gin"
	goJWT "github.com/golang-jwt/jwt/v4"
)

const BearerSchema = "Bearer "

// Auth is a middleware that checks if the request has a valid JWT token and then authorizes the request
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token string from the header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || len(authHeader) <= len(BearerSchema)+1 {
			err := errors.TokenError.New("missing auth token")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			return
		}

		// Cut off the "Bearer " part of the header
		tokenString := authHeader[len(BearerSchema):]

		// Parse and validate the token
		token, err := jwt.Service().ValidateToken(tokenString)
		if err == nil && token.Valid {
			claims := token.Claims.(goJWT.MapClaims)
			uid := claims["uid"].(string)
			log.ApiLogger.WithField("user", claims["uid"].(string)).Info("authorizing user")
			if !db.IsUserEnabled(uid) {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "user is not allowed to use the API"})
				return
			}
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		}
	}
}
