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

	route.Get("/user/detail", middleware.Authorization("all"), controller.GetUserDetail)
	route.Post("/user/update", middleware.Authorization("all"), controller.UpdateUserDetail)

	route.Get("/books", controller.GetBooks)
	route.Get("/book/:id<int>", controller.GetBookById)
	route.Post("/book/insert", middleware.Authorization("admin", "staff"), controller.InsertBook)
	route.Put("/book/update/:id", middleware.Authorization("admin", "staff"), controller.UpdateBook)
	route.Delete("/book/delete/:id", middleware.Authorization("admin", "staff"), controller.DeleteBook)

	route.Get("/book/copies/:id", middleware.Protected(), controller.GetBookCopy)
	route.Post("book/copies/insert", middleware.Protected(), controller.InsertBookCopy)

	route.Get("/book/rent", controller.GetBookRent)
	route.Post("/book/rent/insert", controller.InsertBookRent)
	route.Post("/book/rent/return", controller.ReturnBook)
}
