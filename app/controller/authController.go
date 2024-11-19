package controller

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/IipulI/percobaan-gofiber/app/repository"
	"github.com/IipulI/percobaan-gofiber/database"
	"github.com/gofiber/fiber/v2"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

func Login(c *fiber.Ctx) error {
	payload := struct {
		User     string `json:"user"`
		Password string `json:"password"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	newRepo := repository.NewUserRepository(database.GetDB())
	ctx := context.Background()

	user, err := newRepo.Login(ctx, payload.User, payload.Password)
	if err != nil {
		return err
	}

	// JWT
	t := jwt.New()
	t.Set(jwt.SubjectKey, `localhost:5000`)
	t.Set(jwt.AudienceKey, payload.User)
	t.Set(jwt.IssuedAtKey, time.Now())

	// Signing a token (using raw rsa.PrivateKey)
	signed, err := jwt.Sign(t, jwt.WithKey(jwa.HS256, []byte(os.Getenv("JWT_SECRET_KEY"))))
	if err != nil {
		log.Printf("failed to sign token: %s", err)
		return err
	}

	return c.JSON(fiber.Map{
		"status":  200,
		"message": "success",
		"data": fiber.Map{
			"username": user.Username,
			"email":    user.Email,
			"token":    string(signed),
		},
	})
}
