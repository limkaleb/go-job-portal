package controllers

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/limkaleb/go-job-portal/database"
	"github.com/limkaleb/go-job-portal/models"
)

// 	PostJob	Post job for employer
//	@Summary		Post job for employer
//	@Tags			Employer
//	@Description	This endpoint is used for employer to post a new job. Employer must be authenticated to perform this action.
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.PostJobRequest	true	"Post job request for employer"
//	@Success		201		{object}	models.PostJobResponse
//	@Failure		401		{string}	message	"Unauthenticated"
//	@Failure		404		{string}	message	"User not found"
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

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"data": fiber.Map{"job": newJob}})
}

// 	GetJobs	Get all jobs
//	@Summary		Get all jobs
//	@Tags			Talent
//	@Description	This endpoint is used for talent to fetch all jobs posted. No authentication needed to perform this action.
//	@Produce		json
//	@Success		200	{array}	models.Job
//	@Router			/api/jobs [get]
func GetJobs(c *fiber.Ctx) error {
	var jobs []models.Job

	database.DB.Find(&jobs)

	return c.JSON(fiber.Map{"data": fiber.Map{"jobs": jobs}})
}

// 	GetJobsByEmployer	Employer get jobs
//	@Summary		Employer get jobs
//	@Tags			Employer
//	@Description	This endpoint is used for employer to fetch all jobs he/she posted. Authentication needed to perform this action.
//	@Produce		json
//	@Success		200	{array}		models.Job
//	@Failure		401	{string}	message	"Unauthenticated"
//	@Failure		404	{string}	message	"User not found"
//	@Router			/api/employer/jobs [get]
func GetJobsByEmployer(c *fiber.Ctx) error {
	user := c.Locals("user").(models.Employer)

	var jobs []models.Job

	database.DB.Preload("Applications").Where("employer_id = ?", user.ID).Find(&jobs)

	return c.JSON(fiber.Map{"data": fiber.Map{"jobs": jobs}})
}


// TODO: Update job

// TODO: Delete job
