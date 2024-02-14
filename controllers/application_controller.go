package controllers

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/limkaleb/go-job-portal/database"
	"github.com/limkaleb/go-job-portal/models"
)

func ApplyJob(c *fiber.Ctx) error {
	jobId, _ := strconv.Atoi(c.Params("job_id"))
	user := c.Locals("user").(models.Talent)

	now := time.Now()
	newApplication := models.Application{
		Status: models.Applied,
		SubmissionDate: now,
		TalentID: user.ID,
		JobID: uint(jobId),
		CreatedAt: now,
		UpdatedAt: now,
	}
	database.DB.Create(&newApplication)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"application": newApplication}})
}

func GetApplications(c *fiber.Ctx) error {
	user := c.Locals("user").(models.Talent)

	var applications []models.Application

	database.DB.Where("talent_id = ?", user.ID).Find(&applications)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"applications": applications}})
}

func GetApplicationsById(c *fiber.Ctx) error {
	user := c.Locals("user").(models.Talent)
	appId, _ := strconv.Atoi(c.Params("id"))

	var application models.Application

	database.DB.Where("id = ? AND talent_id = ?", appId, user.ID).Find(&application)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"application": application}})
}
