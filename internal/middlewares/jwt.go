package middlewares

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/n0067h/gofiber-auth-server/internal/apperror"
	"github.com/n0067h/gofiber-auth-server/internal/config"
)

func JWTGuard() fiber.Handler {
	return jwtware.New(jwtware.Config{
		TokenLookup: "header:Authorization",
		AuthScheme:  "Bearer",
		SigningKey:  jwtware.SigningKey{Key: []byte(config.Get().JWTSecret)},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.
				Status(fiber.StatusUnauthorized).
				JSON(fiber.Map{
					"error": apperror.ErrUnauthorized.Error(),
				})
		},
	})
}
