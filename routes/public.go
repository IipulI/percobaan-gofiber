package routes

import (
	"github.com/IipulI/percobaan-gofiber/app/controller"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	route.Get("/books", controller.GetBooks)
	route.Get("/book/:id<int>", controller.GetBookById)
}
