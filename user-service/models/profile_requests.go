package models

import (
	"github.com/google/uuid"
	"time"
)

type CreateProfileRequest struct {
	UserID               uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	Name                 string    `json:"name" gorm:"name"`
	Bio                  string    `json:"bio" gorm:"bio"`
	AvatarURL            string    `json:"avatar_url" gorm:"avatar_url"`
	City                 string    `json:"city" gorm:"city"`
	LookingFor           []string  `json:"looking_for" gorm:"looking_for"`
	Skills               []string  `json:"skills" gorm:"skills"`
	AvailableForProjects bool      `json:"available_for_projects" gorm:"available_for_projects"`
	CreatedAt            time.Time `json:"created_at" gorm:"created_at"`
}
