package middleware

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(500).JSON(fiber.Map{
				"error": "No authorization token",
			})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			fmt.Println(parts[0], len(parts))

			return c.Status(500).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		tokenString := parts[1]
		token, err := jwt.Parse([]byte(tokenString), jwt.WithKey(jwa.HS256, []byte(os.Getenv("JWT_SECRET_KEY"))))
		if err != nil {
			fmt.Print(err)

			return c.Status(500).JSON(fiber.Map{
				"error": "Cannot parse token",
			})
		}

		exp := token.Expiration()
		if exp.Format("2006-01-02") == "0001-01-01" {
			return c.Status(500).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		if time.Now().After(exp) {
			return c.Status(500).JSON(fiber.Map{
				"error": "Token expired",
			})
		}

		resp, ok := token.Get("role")
		if !ok {
			return c.Status(500).JSON(fiber.Map{
				"error": "No claim role found",
			})
		}

		c.Locals("role", resp)
		c.Locals("username", token.Audience()[0])

		return c.Next()
	}
}
