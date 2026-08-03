package models

import (
	"github.com/google/uuid"
	"time"
)

type EmailVerifyToken struct {
	ID               uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID           uuid.UUID `gorm:"type:uuid;not null"`
	EmailVerifyToken string    `gorm:"not null;uniqueIndex"`
	ExpiresAt        time.Time `gorm:"not null"`
	CreatedAt        time.Time
}
