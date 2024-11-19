package controller

import (
	"context"

	"github.com/IipulI/percobaan-gofiber/app/model"
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

func Insert(c *fiber.Ctx) error {
	book := &model.Book{}

	if err := c.BodyParser(book); err != nil {
		panic(err)
	}

	bookRepo := repository.NewBookRepository(database.GetDB())
	ctx := context.Background()

	_, err := bookRepo.Insert(ctx, *book)
	if err != nil {
		panic(err)
	}

	return c.JSON(fiber.Map{
		"status":  201,
		"message": "success",
		"data":    nil,
	})
}

func Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		panic(err)
	}

	bookRepo := repository.NewBookRepository(database.GetDB())
	ctx := context.Background()

	_, err = bookRepo.FindById(ctx, id)
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  422,
			"message": "failed",
			"data":    "",
		})
	}

	book := &model.Book{}
	if err := c.BodyParser(book); err != nil {
		panic(err)
	}

	book.Id = int32(id)

	if _, err := bookRepo.Update(ctx, id, book); err != nil {
		return err
	}

	dbBook, err := bookRepo.FindById(ctx, id)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data":    dbBook,
	})
}

func Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		panic(err)
	}

	bookRepo := repository.NewBookRepository(database.GetDB())
	ctx := context.Background()

	_, err = bookRepo.FindById(ctx, id)
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  422,
			"message": "failed",
			"data":    "",
		})
	}

	msg, err := bookRepo.Delete(ctx, id)
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  422,
			"message": msg,
			"data":    "",
		})
	}

	return c.JSON(fiber.Map{
		"status":  200,
		"message": "deleted successfully",
		"data":    "",
	})
}
