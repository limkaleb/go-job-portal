package database

import (
	"log"
	"os"

	"github.com/limkaleb/go-job-portal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	dsn := "host=localhost user=postgres password=123 dbname=bosshire port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database! \n", err.Error())
		os.Exit(1)
	}

	log.Println("Running Migrations")
	DB.AutoMigrate(&models.Employer{}, &models.Talent{})
	
	log.Println("Connected Successfully to the Database")
}