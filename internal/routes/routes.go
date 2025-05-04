package routes

import (
	"github.com/gofiber/fiber/v2"
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
