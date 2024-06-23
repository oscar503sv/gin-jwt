package models

import (
	"time"

	"gorm.io/gorm"
)

// Modelo para el token invalidado
type BlacklistedToken struct {
	gorm.Model
	ID            uint `gorm:"primaryKey"`
	Token         string
	BlacklistedAt time.Time
}
