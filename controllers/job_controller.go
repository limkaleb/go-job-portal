package controllers

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/limkaleb/go-job-portal/database"
	"github.com/limkaleb/go-job-portal/models"
)

// 	PostJob	Employer post job
//	@Summary		Employer post job
//	@Tags			Employer
//	@Description	Employer post new job
//	@Router			/api/job [post]
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

// 	GetJobs	Get all jobs
//	@Summary		Get all jobs
//	@Tags			Talent
//	@Description	Get all jobs
//	@Router			/api/jobs [get]
func GetJobs(c *fiber.Ctx) error {
	var jobs []models.Job

	database.DB.Find(&jobs)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"jobs": jobs}})
}

// 	GetJobsByEmployer	Employer get jobs
//	@Summary		Employer get jobs
//	@Tags			Employer
//	@Description	Employer get jobs
//	@Router			/api/employer/jobs [get]
func GetJobsByEmployer(c *fiber.Ctx) error {
	user := c.Locals("user").(models.Employer)

	var jobs []models.Job

	database.DB.Preload("Applications").Where("employer_id = ?", user.ID).Find(&jobs)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"jobs": jobs}})
}


// TODO: Update job

// TODO: Delete job