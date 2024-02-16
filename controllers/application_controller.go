package controllers

import (
	"strconv"
	"time"

	"slices"

	"github.com/gofiber/fiber/v2"
	"github.com/limkaleb/go-job-portal/database"
	"github.com/limkaleb/go-job-portal/models"
)

// ApplyJob	Apply Job
//	@Summary		Apply Job
//	@Tags			Talent
//	@Description	This endpoint is used for talent to apply to a job posted. Need authentication to perform this action.
//	@Accept			json
//	@Produce		json
//	@Param			job_id	path		string	true	"Job id"
//	@Success		201		{object}	models.ApplyJobResponse
//	@Failure		401		{string}	message	"Unauthenticated"
//	@Failure		404		{string}	message	"User not found"
//	@Router			/api/job/{job_id}/apply [post]
func ApplyJob(c *fiber.Ctx) error {
	jobId, _ := strconv.Atoi(c.Params("job_id"))
	user := c.Locals("user").(models.Talent)

	now := time.Now()
	newApplication := models.Application{
		Status:         models.Applied,
		SubmissionDate: now,
		TalentID:       user.ID,
		JobID:          uint(jobId),
		CreatedAt:      now,
		UpdatedAt:      now,
	}
	database.DB.Create(&newApplication)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"data": fiber.Map{"application": newApplication}})
}

// GetApplications	Talent get applications
//	@Summary		Talent get applications
//	@Tags			Talent
//	@Description	This endpoint is used for talent to get all their applications. Need authentication to perform this action.
//	@Produce		json
//	@Success		200	{array}		models.Application
//	@Failure		401	{string}	message	"Unauthenticated"
//	@Failure		404	{string}	message	"User not found"
//	@Router			/api/applications [get]
func GetApplications(c *fiber.Ctx) error {
	user := c.Locals("user").(models.Talent)

	var applications []models.Application

	database.DB.Where("talent_id = ?", user.ID).Find(&applications)

	return c.JSON(fiber.Map{"data": fiber.Map{"applications": applications}})
}

// GetApplicationsById	Talent get application by id
//	@Summary		Talent get application by id
//	@Tags			Talent
//	@Description	This endpoint is used for talent to get single application details. Need authentication to perform this action.
//	@Produce		json
//	@Param			id	path		string	true	"Aplication id"
//	@Success		200	{object}	models.Application
//	@Failure		401	{string}	message	"Unauthenticated"
//	@Failure		404	{string}	message	"User not found"
//	@Router			/api/applications/{id} [get]
func GetApplicationsById(c *fiber.Ctx) error {
	user := c.Locals("user").(models.Talent)
	appId, _ := strconv.Atoi(c.Params("id"))

	var application models.Application

	database.DB.Preload("Job").Where("id = ? AND talent_id = ?", appId, user.ID).Find(&application)

	return c.JSON(fiber.Map{"data": fiber.Map{"application": application}})
}

// GetApplicationsByEmployer	Employer get applications
//	@Summary		Employer get applications
//	@Tags			Employer
//	@Description	This endpoint is used for employer to get all their job's applications. Need authentication to perform this action.
//	@Produce		json
//	@Success		200	{array}		models.Application
//	@Failure		401	{string}	message	"Unauthenticated"
//	@Failure		404	{string}	message	"User not found"
//	@Router			/api/employer/applications [get]
func GetApplicationsByEmployer(c *fiber.Ctx) error {
	user := c.Locals("user").(models.Employer)

	var applications []models.Application
	var jobs []models.Job

	database.DB.Preload("Applications").Where("employer_id = ?", user.ID).Find(&jobs)

	for _, job := range jobs {
		applications = append(applications, job.Applications...)
	}

	return c.JSON(fiber.Map{"data": fiber.Map{"applications": applications}})
}

// EmployerGetApplicationById	Employer get applications by id
//	@Summary		Employer get applications by id
//	@Tags			Employer
//	@Description	This endpoint is used for employer to get single application details. Need authentication to perform this action.
//	@Produce		json
//	@Param			id	path		string	true	"Aplication id"
//	@Success		200	{object}	models.Application
//	@Failure		401	{string}	message	"Unauthenticated"
//	@Failure		404	{string}	message	"User not found"
//	@Router			/api/employer/applications/{id} [get]
func EmployerGetApplicationById(c *fiber.Ctx) error {
	user := c.Locals("user").(models.Employer)
	appId, _ := strconv.Atoi(c.Params("id"))

	var application models.Application

	database.DB.Preload("Job").Preload("Talent").Where("id = ?", appId).Find(&application)

	if application.Job.EmployerID != user.ID {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "application not found",
		})
	}

	return c.JSON(fiber.Map{"data": fiber.Map{"application": application}})
}

// UpdateApplicationByEmployer	Employer update application
//	@Summary		Employer update application
//	@Tags			Employer
//	@Description	This endpoint is used for employer to update application status. Status can be "interview", "accept", or "reject". Need authentication to perform this action.
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string							true	"Aplication id"
//	@Param			request	body		models.UpdateApplicationRequest	true	"Update application request for employer"
//	@Success		200		{object}	models.Application
//	@Failure		400		{string}	message	"Wrong status"
//	@Failure		401		{string}	message	"Unauthenticated"
//	@Failure		404		{string}	message	"User not found"
//	@Router			/api/employer/applications/{id} [put]
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
			"message": "Wrong status",
		})
	}

	database.DB.Where("id = ?", appId).Find(&application)
	application.Status = models.Status(data["status"])
	database.DB.Save(&application)

	return c.JSON(fiber.Map{"data": fiber.Map{"application": application}})
}
