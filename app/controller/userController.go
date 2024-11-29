package controller

import (
	"context"

	"github.com/IipulI/percobaan-gofiber/app/model"
	"github.com/IipulI/percobaan-gofiber/app/repository"
	"github.com/IipulI/percobaan-gofiber/app/utils"
	"github.com/IipulI/percobaan-gofiber/database"
	"github.com/gofiber/fiber/v2"
)

func GetUserDetail(c *fiber.Ctx) error {
	newRepo := repository.NewUserDetailRepository(database.GetDB())
	ctx := context.Background()

	result, err := newRepo.GetUserDetail(ctx, "user")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data":    result,
	})
}

func UpdateUserDetail(c *fiber.Ctx) error {
	username := c.Locals("username").(string)
	paramUserDetail := &model.UserDetail{}

	if err := c.BodyParser(paramUserDetail); err != nil {
		return utils.JsonResponse(c, 400, err.Error(), nil)
	}

	newRepo := repository.NewUserDetailRepository(database.GetDB())
	ctx := context.Background()

	update, err := newRepo.UpdateUserDetail(ctx, username, paramUserDetail)

	if err != nil {
		return utils.JsonResponse(c, 400, err.Error(), nil)
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data": fiber.Map{
			"first_name":      update.FirstName,
			"last_name":       update.LastName,
			"address":         update.Address,
			"phone_number":    update.PhoneNumber,
			"gender":          update.Gender,
			"date_of_birth":   update.DateOfBirth,
			"profile_picture": update.ProfilePicture,
		},
	})
}
