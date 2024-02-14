package controllers

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/limkaleb/go-job-portal/database"
	"github.com/limkaleb/go-job-portal/models"
)

func PostJob(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	user := c.Locals("user").(models.Employer)

	now := time.Now()
	salary, _ := strconv.ParseUint(data["salary"], 10, 64)
	newJob := models.Job{
		Title: data["title"],
		Description: data["description"],
		Salary: uint(salary),
		Location: data["location"],
		EmployerID: user.ID,
		CreatedAt: now,
		UpdatedAt: now,
	}
	database.DB.Create(&newJob)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"job": newJob}})
}

// Get jobs

// Update Job

// Delete Job