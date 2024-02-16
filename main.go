package main

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/limkaleb/go-job-portal/database"
	"github.com/limkaleb/go-job-portal/routes"
)

// @title			Job Portal API
// @version		1.0
// @description	The Simple Job Portal is an API-based platform designed to connect job seekers (talents)
//
//	with employers. It provides two distinct user flows: one for talents to search and apply for
//	jobs, and another for employers to post jobs, review applications, and manage the hiring
//	process.
//
// @termsOfService	http://swagger.io/terms/
func main() {
	database.Connect()
	app := fiber.New()
	routes.Setup(app)

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	app.Use(csrf.New())

	app.Use(logger.New())

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))

	app.Listen(":8000")
}
