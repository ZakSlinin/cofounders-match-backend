package profile_service

import (
	"context"
	"github.com/ZakSlinin/cofounders-match-backend/user-service/models"
	profile_repository "github.com/ZakSlinin/cofounders-match-backend/user-service/profile/profile-repository"
)

type ProfileService interface {
	Create(ctx context.Context, profile *models.Profile) (*models.Profile, error)
	UpdateAvatar(ctx context.Context, userID string, avatarURL string) error
}

type RepoProfileService struct {
	repo profile_repository.ProfileRepository
}

func NewProfileService(repo profile_repository.ProfileRepository) ProfileService {
	return &RepoProfileService{repo: repo}
}

func (service *RepoProfileService) Create(ctx context.Context, profile *models.Profile) (*models.Profile, error) {
	result, err := service.repo.Create(ctx, profile)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *RepoProfileService) UpdateAvatar(ctx context.Context, userID string, avatarURL string) error {
	return service.repo.UpdateAvatar(ctx, userID, avatarURL)
}
