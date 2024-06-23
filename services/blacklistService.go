package services

import (
	"time"

	"github.com/oscar503sv/gin-jwt/config"
	"github.com/oscar503sv/gin-jwt/models"
)

// Añadir un token a la lista negra
func AddToBlacklist(token string) error {
	blacklistedToken := models.BlacklistedToken{
		Token:         token,
		BlacklistedAt: time.Now(),
	}
	return config.DB.Create(&blacklistedToken).Error
}

// Comprobar si un token está en la lista negra
func IsBlacklisted(token string) (bool, error) {
	var count int64
	err := config.DB.Model(&models.BlacklistedToken{}).Where("token = ?", token).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
