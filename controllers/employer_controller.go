package controllers

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/limkaleb/go-job-portal/database"
	"github.com/limkaleb/go-job-portal/models"
	"golang.org/x/crypto/bcrypt"
)

// 	RegisterEmployer	Register Employer
//	@Summary		Register employer
//	@Tags			Employer
//	@Description	Register employer
//	@Router			/api/employer/register [post]
func RegisterEmployer(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	now := time.Now()
	employer := models.Employer{
		Name:      data["name"],
		Email:     data["email"],
		Password:  password,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result := database.DB.Create(&employer)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "Email already exist!"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"employer": employer}})
}

// 	LoginEmployer	Login Employer
//	@Summary		Login employer
//	@Tags			Employer
//	@Description	Login employer
//	@Router			/api/employer/login [post]
func LoginEmployer(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.Employer

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.ID == 0 {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"iss": user.Email,
		"exp": time.Now().Add(time.Hour * 24).Unix(), // 1 day
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Could not login"})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    t,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "Login success",
	})
}

// 	LogoutEmployer	Logout Employer
//	@Summary		Logout employer
//	@Tags			Employer
//	@Description	Get employer data
//	@Router			/api/employer/logout [post]
func LogoutEmployer(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "Logout success",
	})
}

// 	GetEmployer Get Employer
//	@Summary		Get employer data
//	@Tags			Employer
//	@Description	Get employer data
//	@Router			/api/employer [get]
func GetEmployer(c *fiber.Ctx) error {
	user := c.Locals("user")
	return c.JSON(user)
}
