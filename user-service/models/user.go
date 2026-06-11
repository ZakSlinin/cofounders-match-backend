package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Email        string    `gorm:"uniqueIndex;not null"                           json:"email"`
	PasswordHash string    `gorm:"not null"                                       json:"-"`
	Role         string    `gorm:"not null"                                       json:"role"`
}
