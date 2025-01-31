package controller

import (
	"context"

	"github.com/IipulI/percobaan-gofiber/app/model"
	"github.com/IipulI/percobaan-gofiber/app/repository"
	"github.com/IipulI/percobaan-gofiber/app/utils"
	"github.com/IipulI/percobaan-gofiber/database"
	"github.com/gofiber/fiber/v2"
)

func GetBookCopy(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return utils.JsonResponse(c, 400, err.Error(), "")
	}

	newRepo := repository.NewBookCopiesRepository(database.GetDB())
	ctx := context.Background()

	result, err := newRepo.GetBookCopyByBook(ctx, id)
	if err != nil {
		return utils.JsonResponse(c, 400, err.Error(), "")
	}

	return utils.JsonResponse(c, 200, "success", result)
}

func InsertBookCopy(c *fiber.Ctx) error {
	bookCopy := &model.BookCopies{}
	if err := c.BodyParser(bookCopy); err != nil {
		return utils.JsonResponse(c, 400, err.Error(), "")
	}

	bookRepo := repository.NewBookRepository(database.GetDB())
	bookCopyRepo := repository.NewBookCopiesRepository(database.GetDB())
	ctx := context.Background()

	_, err := bookRepo.FindById(ctx, int(bookCopy.BookId))
	if err != nil {
		return utils.JsonResponse(c, 400, err.Error(), "")
	}

	if err = bookCopyRepo.ValidateBookCopyNumber(ctx, *bookCopy); err != nil {
		return utils.JsonResponse(c, 400, err.Error(), "")
	}

	_, err = bookCopyRepo.InsertBookCopy(ctx, *bookCopy)
	if err != nil {
		return utils.JsonResponse(c, 400, err.Error(), "")
	}

	return utils.JsonResponse(c, 200, "success", "")
}
