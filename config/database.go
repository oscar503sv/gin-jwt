package config

import (
	"log"

	"github.com/oscar503sv/gin-jwt/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("user.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	// Migrar el modelo User
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("failed to migrate database: User", err)
	}

	// Migrar el modelo BlacklistedToken
	if err := DB.AutoMigrate(&models.BlacklistedToken{}); err != nil {
		log.Fatal("failed to migrate database: BlacklistedToken", err)
	}
}
