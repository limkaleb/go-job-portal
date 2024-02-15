package controllers

import (
	"strconv"
	"time"

	"slices"

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


func GetApplicationsByEmployer(c *fiber.Ctx) error {
	user := c.Locals("user").(models.Employer)

	var applications []models.Application
	var jobs []models.Job // TODO: fix response

	database.DB.Preload("Applications").Where("employer_id = ?", user.ID).Find(&jobs)

	for _, job := range jobs {
		applications = append(applications, job.Applications...)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"applications": applications}})
}


func UpdateApplicationByEmployer(c *fiber.Ctx) error {
	user := c.Locals("user").(models.Employer)
	appId, _ := strconv.Atoi(c.Params("id"))

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var application models.Application

	database.DB.Preload("Job").Where("id = ?", appId).Find(&application)

	if application.Job.EmployerID != user.ID {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "application not found",
		})
	}

	statuses := []string{string(models.Accept), string(models.Applied), string(models.Interview), string(models.Reject)}
	if !slices.Contains(statuses, data["status"]) {
		c.Status(fiber.ErrBadRequest.Code)
		return c.JSON(fiber.Map{
			"message": "wrong status",
		})
	}

	database.DB.Where("id = ?", appId).Find(&application)
	application.Status = models.Status(data["status"])
	database.DB.Save(&application)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"application": application}})
}