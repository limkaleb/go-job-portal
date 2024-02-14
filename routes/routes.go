package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/limkaleb/go-job-portal/controllers"
	"github.com/limkaleb/go-job-portal/middlewares"
)

func Setup(app *fiber.App) {	
	// Employer Routes
	app.Post("/api/employer/register", controllers.RegisterEmployer)
	app.Post("/api/employer/login", controllers.LoginEmployer)
	app.Get("/api/employer", middlewares.EmployerAuthMiddleware, controllers.GetEmployer)
	app.Post("/api/employer/logout", controllers.LogoutEmployer)

	app.Post("/api/job", middlewares.EmployerAuthMiddleware, controllers.PostJob)

	// Talent Routes
	app.Post("/api/talent/register", controllers.RegisterTalent)
	app.Post("/api/talent/login", controllers.LoginTalent)
	app.Get("/api/talent", middlewares.TalentAuthMiddleware, controllers.GetTalent)
	app.Post("/api/talent/logout", controllers.LogoutTalent)

	app.Post("/api/job/:job_id/apply", middlewares.TalentAuthMiddleware, controllers.ApplyJob)
	app.Get("/api/applications", middlewares.TalentAuthMiddleware, controllers.GetApplications)
	app.Get("/api/applications/:id", middlewares.TalentAuthMiddleware, controllers.GetApplicationsById)
}