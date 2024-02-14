package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/limkaleb/go-job-portal/database"
	"github.com/limkaleb/go-job-portal/routes"
)

func main() {
	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8000")
}
