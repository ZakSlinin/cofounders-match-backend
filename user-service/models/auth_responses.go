package models

import "github.com/google/uuid"

type AuthResponse struct {
	ID           uuid.UUID `gorm:"primary_key;type:uuid" json:"id"`
	Email        string    `gorm:"column:email;type:varchar(255)" json:"email"`
	Role         string    `gorm:"column:role;type:varchar(50)" json:"role"`
	RefreshToken string    `gorm:"column:refresh_token" json:"refresh_token"`
	AccessToken  string    `gorm:"column:access_token" json:"access_token"`
}
