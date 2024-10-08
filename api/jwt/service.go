package jwt

import (
	"time"

	"github.com/buonotti/apisense/errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

type JwtService interface {
	GenerateToken(userId string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
	RefreshToken(token string, userId string) (string, error)
}

type authCustomClaims struct {
	Uid string `json:"uid"`
	jwt.RegisteredClaims
}

type jwtServices struct {
	secretKey string
	issure    string
}

func Service() JwtService {
	return &jwtServices{
		secretKey: viper.GetString("APISENSE_SIGNING_KEY"),
		issure:    "Apisense WEB-API",
	}
}

func (service *jwtServices) GenerateToken(uid string) (string, error) {
	claims := &authCustomClaims{
		uid,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 48)),
			Issuer:    service.issure,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(service.secretKey))
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, errors.TokenError.New("invalid token %s", token.Header["alg"])
		}
		return []byte(service.secretKey), nil
	})
}

func (service *jwtServices) RefreshToken(token string, userId string) (string, error) {
	if _, err := service.ValidateToken(token); err != nil {
		return "", err
	}

	return service.GenerateToken(userId)
}
