package jwt

import (
	"time"

	"github.com/buonotti/apisense/v2/errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

// Issuer holds the issuer of the token
const Issuer string = "Apisense WEB-API"

// authCustomClaims is the data held by the token
type authCustomClaims struct {
	Uid string `json:"uid"`
	jwt.RegisteredClaims
}

// GenerateToken generates a new token for the given user
func GenerateToken(uid string) (string, error) {
	secretKey := viper.GetString("api.signing_key")
	claims := &authCustomClaims{
		uid,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 48)),
			Issuer:    Issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secretKey))
}

// ValidateToken validates if the given token is valid. Returns the decoded token
func ValidateToken(encodedToken string) (*jwt.Token, error) {
	secretKey := viper.GetString("api.signing_key")
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, errors.TokenError.New("invalid token %s", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
}

// RefreshToken takes a valid token and generates a new one from it
func RefreshToken(token string, userId string) (string, error) {
	if _, err := ValidateToken(token); err != nil {
		return "", err
	}

	return GenerateToken(userId)
}
