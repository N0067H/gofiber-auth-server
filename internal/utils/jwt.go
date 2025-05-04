package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/n0067h/gofiber-auth-server/internal/apperror"
	"github.com/n0067h/gofiber-auth-server/internal/config"
)

func GenerateJWT(id int, name string) (string, error) {
	claims := jwt.MapClaims{
		"sub":  id,
		"name": name,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Minute * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(config.Get().JWTSecret))
	if err != nil {
		return t, apperror.ErrInternalServerError
	}

	return t, nil
}
