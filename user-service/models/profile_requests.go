package models

type CreateProfileRequest struct {
	Name                 string   `json:"name"                  binding:"required"`
	Bio                  string   `json:"bio"                   binding:"required"`
	City                 string   `json:"city"`
	LookingFor           []string `json:"looking_for"`
	Skills               []string `json:"skills"`
	AvailableForProjects bool     `json:"available_for_projects"`
}
