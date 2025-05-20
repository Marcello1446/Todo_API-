package database

import (
	"Todo/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	connStr := "host=localhost user=postgres dbname=postgres password=Admin port=5432 sslmode=disable"

	var err error
	DB, _ = gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect a database!")
	}

	err = DB.AutoMigrate(&models.Task{})

	if err != nil {
		log.Fatal("Failed to find structs!")
	}
}
