package models

import "github.com/google/uuid"

//name
//username
//role
//bio
//avatar
//city
//looking_for
//skills
//available_for_projects

type Profile struct {
	ID   uuid.UUID `gorm:"primary_key" json:"id"`
	Name string    `json:"name" gorm:"name"`
}
