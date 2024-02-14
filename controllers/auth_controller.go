package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/limkaleb/go-job-portal/models"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	employer := models.Employer{
		Name: data["name"],
		Email: data["email"],
		Password: password,
	}

	return c.JSON(employer)
}
