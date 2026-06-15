package profile_service

import (
	"context"
	"github.com/ZakSlinin/cofounders-match-backend/user-service/models"
	profile_repository "github.com/ZakSlinin/cofounders-match-backend/user-service/profile/profile-repository"
)

type ProfileService struct {
	repo profile_repository.ProfileRepository
}

func NewProfileService(repo profile_repository.ProfileRepository) *ProfileService {
	return &ProfileService{repo: repo}
}

func (service *ProfileService) Create(ctx context.Context, profile *models.Profile) (*models.Profile, error) {
	result, err := service.repo.Create(ctx, profile)
	if err != nil {
		return nil, err
	}

	return result, nil
}
