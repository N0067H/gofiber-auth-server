package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/n0067h/gofiber-auth-server/internal/config"
	"github.com/n0067h/gofiber-auth-server/internal/db"
	"github.com/n0067h/gofiber-auth-server/internal/routes"
)

func main() {
	config.Init()
	if err := db.Init(config.Get()); err != nil {
		log.Fatalf("failed to initialize database connection: %v", err)
	}

	if err := db.Migrate(db.Get()); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("failed to close database: %v", err)
		}
	}()

	app := fiber.New()

	routes.SetupRoutes(app, db.Get())

	log.Fatal(app.Listen(":3000"))
}
