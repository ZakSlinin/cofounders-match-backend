package models

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"time"
)

type Profile struct {
	ID                   uuid.UUID      `gorm:"type:uuid;primary_key"        json:"id"`
	UserID               uuid.UUID      `gorm:"type:uuid;not null"           json:"user_id"`
	Name                 string         `gorm:"column:name"                  json:"name"`
	Bio                  string         `gorm:"column:bio"                   json:"bio"`
	AvatarURL            string         `gorm:"column:avatar_url"            json:"avatar_url"`
	City                 string         `gorm:"column:city"                  json:"city"`
	LookingFor           pq.StringArray `gorm:"column:looking_for;type:text[]" json:"looking_for"`
	Skills               pq.StringArray `gorm:"column:skills;type:text[]"    json:"skills"`
	AvailableForProjects bool           `gorm:"column:available_for_projects" json:"available_for_projects"`
	CreatedAt            time.Time      `gorm:"column:created_at"            json:"created_at"`
}

func (Profile) TableName() string {
	return "profiles"
}
