package jwt

import (
	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v4"
)

func ExtractUidFromToken(token *jwt.Token) string {
	return token.Claims.(jwt.MapClaims)["uid"].(string)
}

func ExtractUidFromHeader(c *fiber.Ctx) string {
	const BEARER_SCHEMA = "Bearer "
	authHeader := c.Get("Authorization")
	tokenString := authHeader[len(BEARER_SCHEMA):]
	token, err := Service().ValidateToken(tokenString)
	if err != nil {
		return ""
	}
	claims := token.Claims.(jwt.MapClaims)
	return claims["uid"].(string)
}
