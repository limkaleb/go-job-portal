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

// RegisterEmployer Register an employer
//	@Summary		Register an employer
//	@Tags			Employer
//	@Description	Register as an employer
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.EmployerRegisterRequest	true	"Employer register request"
//	@Success		201		{object}	models.EmployerRegisterResponse
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
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"message": "Email already exist!"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"message": result.Error.Error()})
	}

	employer.Password = []byte{}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"data": fiber.Map{"employer": employer}})
}

// LoginEmployer	Login an employer
//	@Summary		Login an employer
//	@Tags			Employer
//	@Description	Login as an employer then generate JWT token, and store it inside cookies.
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.EmployerLoginRequest	true	"Employer login request"
//	@Success		200		{string}	status						ok
//	@Failure		400		{string}	message						"Input an email"
//	@Failure		401		{string}	message						"Incorrect password"
//	@Failure		404		{string}	message						"User not found"
//	@Failure		500		{string}	message						"Could not login"
//	@Router			/api/employer/login [post]
func LoginEmployer(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["email"] == "" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Input an email",
		})
	}

	var user models.Employer
	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Incorrect password",
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
//	@Description	Logout employer
//	@Produce		json
//	@Success		200	{string}	status	ok
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
//	@Summary		Get current employer data
//	@Tags			Employer
//	@Description	Get current employer data
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.GetEmployerResponse
//	@Failure		401	{string}	message	"Unauthenticated"
//	@Failure		404	{string}	message	"User not found"
//	@Router			/api/employer [get]
func GetEmployer(c *fiber.Ctx) error {
	user := c.Locals("user")
	return c.JSON(user)
}
