package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/n0067h/gofiber-auth-server/ent"
	"github.com/n0067h/gofiber-auth-server/internal/handlers"
)

func SetupRoutes(app *fiber.App, client *ent.Client) {
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong")
	})
	app.Post("/register", handlers.Register(client))
	app.Post("/login", handlers.Login(client))
}

func SetupRestrictedRoutes(app *fiber.App, client *ent.Client) {
	app.Get("/restricted", func(c *fiber.Ctx) error {
		u := c.Locals("user").(*jwt.Token)
		claims := u.Claims.(jwt.MapClaims)
		name := claims["name"].(string)
		return c.SendString("Hello " + name)
	})
}
