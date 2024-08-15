package main

import (
	"os"

	"github.com/IipulI/percobaan-gofiber/config"
	"github.com/IipulI/percobaan-gofiber/database"
	"github.com/IipulI/percobaan-gofiber/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadConfig()

	if err := database.Connection(); err != nil {
		panic("Failed to connect database")
	}

	app := fiber.New()

	routes.PublicRoutes(app)

	app.Listen(os.Getenv("APP_PORT"))
}
