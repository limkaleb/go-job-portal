package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/limkaleb/go-job-portal/controllers"
)

func Setup(app *fiber.App) {
	// Employer Routes
	app.Post("/api/employer/register", controllers.RegisterEmployer)
	app.Post("/api/employer/login", controllers.LoginEmployer)
	app.Get("/api/employer", controllers.GetEmployer)
	app.Post("/api/employer/logout", controllers.LogoutEmployer)
}