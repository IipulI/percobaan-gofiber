package controller

import (
	"context"

	"github.com/IipulI/percobaan-gofiber/app/repository"
	"github.com/IipulI/percobaan-gofiber/database"

	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	newRepo := repository.NewBookRepository(database.GetDB())
	ctx := context.Background()

	result, err := newRepo.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	return c.JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data":    result,
	})
}

func GetBookById(c *fiber.Ctx) error {
	idBook, err := c.ParamsInt("id")
	if err != nil {
		panic(err)
	}

	newRepo := repository.NewBookRepository(database.GetDB())
	ctx := context.Background()

	result, err := newRepo.FindById(ctx, idBook)
	if err != nil {
		panic(err)
	}

	return c.JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data":    result,
	})
}
