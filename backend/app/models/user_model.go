package models

import (
	"time"

	// "github.com/satori/go.uuid"
)

// User struct to describe User object.
type User struct {
	ID           uint `gorm:"primaryKey"`
	Username     string    `gorm:"unique;not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Email        string `gorm:"unique;not null"`
	PasswordHash string
	UserStatus   int
	UserRole     string
}
