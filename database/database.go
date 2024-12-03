package database

import (
	"go-blog/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// db is the internal database connection
var db *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = database.AutoMigrate(&models.UsersModel{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Store the connection in the db variable
	db = database
}

// GetDB returns the global db object
func GetDB() *gorm.DB {
	return db
}
