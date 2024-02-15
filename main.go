package main

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/limkaleb/go-job-portal/database"
	"github.com/limkaleb/go-job-portal/routes"
)

//	@title			Job Portal API
//	@version		1.0
//	@description	This is a very simple job portal API
//	@termsOfService	http://swagger.io/terms/
func main() {
	database.Connect()
	app := fiber.New()
	routes.Setup(app)

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	app.Use(logger.New())

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))

	app.Listen(":8000")
}
