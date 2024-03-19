package db

import "example.com/bebelino/models"

func MigrateDatabase() {
	DB.AutoMigrate(&models.Student{})
}