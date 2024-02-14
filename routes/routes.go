package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/limkaleb/go-job-portal/controllers"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
}