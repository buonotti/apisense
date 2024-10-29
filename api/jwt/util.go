package jwt

import (
	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v4"
)

// ExtractUidFromToken extracts the uid from a given token
func ExtractUidFromToken(token *jwt.Token) string {
	return token.Claims.(jwt.MapClaims)["uid"].(string)
}

// ExtractUidFromHeader extracts the uid from the token in the Authorization header
func ExtractUidFromHeader(c *fiber.Ctx) string {
	const bearerSchema = "Bearer "
	authHeader := c.Get("Authorization")
	tokenString := authHeader[len(bearerSchema):]
	token, err := ValidateToken(tokenString)
	if err != nil {
		return ""
	}
	claims := token.Claims.(jwt.MapClaims)
	return claims["uid"].(string)
}
