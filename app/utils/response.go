package utils

import "github.com/gofiber/fiber/v2"

func JsonResponse(c *fiber.Ctx, status int, message string, data interface{}) error {
	if status >= 200 && status <= 300 {
		return c.Status(status).JSON(fiber.Map{
			"status":  status,
			"message": message,
			"data":    data,
		})
	} else {
		return c.Status(status).JSON(fiber.Map{
			"status": status,
			"error":  message,
		})
	}

}
