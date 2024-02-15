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

// 	RegisterTalent	Register Talent
//	@Summary		Register talent
//	@Tags			Talent
//	@Description	Register talent
//	@Router			/api/talent/register [post]
func RegisterTalent(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	now := time.Now()
	talent := models.Talent{
		Name: data["name"],
		Email: data["email"],
		Password: password,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result := database.DB.Create(&talent)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "Email already exist!"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"talent": talent}})
}

// 	LoginTalent	Login Talent
//	@Summary		Login talent
//	@Tags			Talent
//	@Description	Login talent
//	@Router			/api/talent/login [post]
func LoginTalent(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.Talent

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
		"iss":  user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // 1 day
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

// 	LogoutTalent	Logout Talent
//	@Summary		Logout talent
//	@Tags			Talent
//	@Description	Get talent data
//	@Router			/api/talent/logout [post]
func LogoutTalent(c *fiber.Ctx) error {
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

// 	GetTalent Get Talent
//	@Summary		Get talent data
//	@Tags			Talent
//	@Description	Get talent data
//	@Router			/api/talent [get]
func GetTalent(c *fiber.Ctx) error {
	user := c.Locals("user")
	return c.JSON(user)
}