package database

import (
	"log"

	"github.com/bsorawit1234/e-commerce-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("ecommerce.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = database

	err = DB.AutoMigrate(&models.Product{}, &models.Order{}, &models.OrderItem{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
