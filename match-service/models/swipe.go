package models

import "time"

type Swipe struct {
	ID        string `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	FromUser  string `gorm:"type:uuid;not null"`
	ToUser    string `gorm:"type:uuid;not null"`
	CreatedAt time.Time
}
