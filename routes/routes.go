package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/limkaleb/go-job-portal/controllers"
	_ "github.com/limkaleb/go-job-portal/docs"
	"github.com/limkaleb/go-job-portal/middlewares"
)

func Setup(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "Welcome to Job Portal API",
		})
	})

	// Employer routes
	app.Post("/api/employer/register", controllers.RegisterEmployer)
	app.Post("/api/employer/login", controllers.LoginEmployer)
	app.Get("/api/employer", middlewares.EmployerAuthMiddleware, controllers.GetEmployer)
	app.Post("/api/employer/logout", controllers.LogoutEmployer)

	app.Post("/api/job", middlewares.EmployerAuthMiddleware, controllers.PostJob)
	app.Get("/api/employer/jobs", middlewares.EmployerAuthMiddleware, controllers.GetJobsByEmployer)
	app.Get("/api/employer/applications", middlewares.EmployerAuthMiddleware, controllers.GetApplicationsByEmployer)
	app.Put("/api/employer/applications/:id", middlewares.EmployerAuthMiddleware, controllers.UpdateApplicationByEmployer)
	app.Get("/api/employer/applications/:id", middlewares.EmployerAuthMiddleware, controllers.EmployerGetApplicationById)

	// Talent routes
	app.Post("/api/talent/register", controllers.RegisterTalent)
	app.Post("/api/talent/login", controllers.LoginTalent)
	app.Get("/api/talent", middlewares.TalentAuthMiddleware, controllers.GetTalent)
	app.Post("/api/talent/logout", controllers.LogoutTalent)

	app.Post("/api/job/:job_id/apply", middlewares.TalentAuthMiddleware, controllers.ApplyJob)
	app.Get("/api/applications", middlewares.TalentAuthMiddleware, controllers.GetApplications)
	app.Get("/api/applications/:id", middlewares.TalentAuthMiddleware, controllers.GetApplicationsById)

	// Public routes
	app.Get("/api/jobs", controllers.GetJobs)

	// Swagger route
	app.Get("/swagger/*", swagger.HandlerDefault)

}
