package database

import (
	"go-job-portal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=123 dbname=bosshire port=5432 sslmode=disable"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	conn.AutoMigrate(&models.Employer{})
	conn.AutoMigrate(&models.Talent{})
}