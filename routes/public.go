package routes

import (
	"github.com/IipulI/percobaan-gofiber/app/controller"
	"github.com/IipulI/percobaan-gofiber/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {
	api := a.Group("/api")

	api.Post("/login", controller.Login)

	route := a.Group("/api/v1", middleware.Protected())

	route.Get("/user/detail", middleware.Authorization("admin", "staff", "user"), controller.GetUserDetail)
	route.Post("/user/update", middleware.Authorization("admin", "staff", "user"), controller.UpdateUserDetail)

	route.Get("/books", controller.GetBooks)
	route.Get("/book/:id<int>", controller.GetBookById)
	route.Post("/book/insert", middleware.Authorization("admin", "staff"), controller.Insert)
	route.Put("/book/update/:id", middleware.Authorization("admin", "staff"), controller.Update)
	route.Delete("/book/delete/:id", middleware.Authorization("admin", "staff"), controller.Delete)
}
