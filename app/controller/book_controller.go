package controller

import (
	"context"
	"time"

	"github.com/IipulI/percobaan-gofiber/app/model"
	"github.com/IipulI/percobaan-gofiber/app/repository"
	"github.com/IipulI/percobaan-gofiber/app/utils"
	"github.com/IipulI/percobaan-gofiber/database"

	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	newRepo := repository.NewBookRepository(database.GetDB())
	ctx := context.Background()

	result, err := newRepo.FindAll(ctx)
	if err != nil {
		return utils.JsonResponse(c, 400, err.Error(), "")
	}

	return utils.JsonResponse(c, 200, "success", result)
}

func GetBookById(c *fiber.Ctx) error {
	idBook, err := c.ParamsInt("id")
	if err != nil {
		return utils.JsonResponse(c, 400, err.Error(), "")
	}

	newRepo := repository.NewBookRepository(database.GetDB())
	ctx := context.Background()

	result, err := newRepo.FindById(ctx, idBook)
	if err != nil {
		return utils.JsonResponse(c, 400, err.Error(), "")
	}

	return utils.JsonResponse(c, 200, "success", result)
}

func InsertBook(c *fiber.Ctx) error {
	book := &model.Book{}

	// parsing data dari body request ke model book
	if err := c.BodyParser(book); err != nil {
		return utils.JsonResponse(c, 400, err.Error(), "")
	}
	// book.CreatedAt = utils.NewCustomDateTime(time.Now())

	bookRepo := repository.NewBookRepository(database.GetDB())
	ctx := context.Background()

	_, err := bookRepo.Insert(ctx, *book)
	if err != nil {
		return utils.JsonResponse(c, 400, err.Error(), "")
	}

	return utils.JsonResponse(c, 200, "success", nil)
}

func UpdateBook(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return utils.JsonResponse(c, 400, "invalid ID parameter", nil)
	}

	bookRepo := repository.NewBookRepository(database.GetDB())
	ctx := context.Background()

	// Validasi apakah ID ada
	_, err = bookRepo.FindById(ctx, id)
	if err != nil {
		return utils.JsonResponse(c, 404, "Book not found", nil)
	}

	// Parse body request
	book := &model.Book{}
	if err := c.BodyParser(book); err != nil {
		return utils.JsonResponse(c, 400, "invalid request body", nil)
	}
	book.UpdatedAt = utils.NewCustomDateTime(time.Now())

	// Validasi data
	if book.Title == "" {
		return utils.JsonResponse(c, 400, "Book name is required", nil)
	}

	// Update book
	updatedBook, err := bookRepo.Update(ctx, id, book)
	if err != nil {
		return utils.JsonResponse(c, 500, err.Error(), nil)
	}

	return utils.JsonResponse(c, 200, "success", updatedBook)
}

func DeleteBook(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		utils.JsonResponse(c, 400, err.Error(), "")
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
	book.DeletedAt = utils.NewCustomDateTime(time.Now())
	msg, err := bookRepo.Delete(ctx, id, book)
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
