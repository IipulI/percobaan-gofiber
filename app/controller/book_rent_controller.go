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

func GetBookRent(c *fiber.Ctx) error {
	newRepo := repository.NewBookRentRepository(database.GetDB())
	ctx := context.Background()

	result, err := newRepo.GetBookRent(ctx)
	if err != nil {
		utils.JsonResponse(c, 400, err.Error(), "")
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data":    result,
	})
}

func InsertBookRent(c *fiber.Ctx) error {
	bookRent := &model.BookRent{}
	if err := c.BodyParser(bookRent); err != nil {
		return utils.JsonResponse(c, 400, err.Error(), "")
	}

	db := database.GetDB()
	bookRentRepo := repository.NewBookRentRepository(db)
	bookCopyRepo := repository.NewBookCopiesRepository(db)
	ctx := context.Background()

	bookCopy, err := bookCopyRepo.GetBookCopyById(ctx, int(bookRent.BookCopyId))
	if err != nil {
		return utils.JsonResponse(c, 400, err.Error(), "")
	}

	if err = bookCopyRepo.ValidateBookCopyStatus(ctx, int(bookRent.BookCopyId)); err != nil {
		return utils.JsonResponse(c, 400, err.Error(), "")
	}

	tx, err := db.Begin()
	if err != nil {
		return utils.JsonResponse(c, 400, err.Error(), "")
	}
	defer tx.Rollback()

	_, err = bookRentRepo.InsertBookRent(ctx, tx, *bookRent)
	if err != nil {
		return utils.JsonResponse(c, 400, err.Error(), "")
	}

	bookCopy.Status = "rent"

	_, err = bookCopyRepo.UpdateBookCopy(ctx, tx, &bookCopy)
	if err != nil {
		return utils.JsonResponse(c, 400, err.Error(), "")
	}

	if err := tx.Commit(); err != nil {
		return utils.JsonResponse(c, 400, err.Error(), "")
	}

	return utils.JsonResponse(c, 200, "success", "")
}

func ReturnBook(c *fiber.Ctx) error {
	bookRent := &model.BookRent{}
	if err := c.BodyParser(bookRent); err != nil {
		return utils.JsonResponse(c, 400, err.Error(), "")
	}
	bookRent.UpdatedAt = utils.NewCustomDateTime(time.Now())

	db := database.GetDB()
	bookRentRepo := repository.NewBookRentRepository(db)
	bookCopyRepo := repository.NewBookCopiesRepository(db)
	ctx := context.Background()

	databaseBookRent, err := bookRentRepo.GetBookRentById(ctx, int(bookRent.Id))
	if err != nil {
		return utils.JsonResponse(c, 400, err.Error(), "")
	}

	bookCopy, err := bookCopyRepo.GetBookCopyById(ctx, int(databaseBookRent.BookCopyId))
	if err != nil {
		return utils.JsonResponse(c, 400, err.Error(), "")
	}

	tx, err := db.Begin()
	if err != nil {
		return utils.JsonResponse(c, 400, err.Error(), "")
	}
	defer tx.Rollback()

	_, err = bookRentRepo.UpdateBookRent(ctx, tx, bookRent)
	if err != nil {
		return utils.JsonResponse(c, 400, err.Error(), "")
	}

	bookCopy.Status = "available"
	_, err = bookCopyRepo.UpdateBookCopy(ctx, tx, &bookCopy)
	if err != nil {
		return utils.JsonResponse(c, 400, err.Error(), "")
	}

	if err := tx.Commit(); err != nil {
		return utils.JsonResponse(c, 400, err.Error(), "")
	}

	return utils.JsonResponse(c, 200, "success", "")
}

func GetBookRentDetail(c *fiber.Ctx) error {

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data":    nil,
	})
}
