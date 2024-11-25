package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func Authorization(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Retrieve the role from context (set by the JWT middleware)
		userRole := c.Locals("role")

		// Ensure the role exists in the context
		if userRole == nil {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Forbidden: No role found",
			})
		}

		// Check if the user's role is one of the allowed roles
		for _, role := range allowedRoles {
			if userRole == role {
				// If valid, proceed with the request
				return c.Next()
			}
		}

		// If no matching role, return forbidden
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Forbidden: You do not have access to this resource",
		})
	}
}
