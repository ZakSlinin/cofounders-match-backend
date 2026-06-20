package profile_service

import (
	"context"
	"github.com/ZakSlinin/cofounders-match-backend/user-service/models"
	profile_repository "github.com/ZakSlinin/cofounders-match-backend/user-service/profile/profile-repository"
	"github.com/google/uuid"
)

type ProfileService interface {
	Create(ctx context.Context, profile *models.Profile) (*models.Profile, error)
	UpdateAvatar(ctx context.Context, userID string, avatarURL string) error
	GetByUserID(ctx context.Context, userID uuid.UUID) (*models.Profile, error)
	UpdateProfile(ctx context.Context, userID uuid.UUID, update *models.UpdateProfileRequest) (*models.Profile, error)
	GetFeed(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*models.Profile, error)
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

func (service *RepoProfileService) GetByUserID(ctx context.Context, userID uuid.UUID) (*models.Profile, error) {
	result, err := service.repo.GetByUserID(ctx, userID)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *RepoProfileService) UpdateProfile(ctx context.Context, userID uuid.UUID, update *models.UpdateProfileRequest) (*models.Profile, error) {
	return service.repo.UpdateProfile(ctx, userID, update)
}

func (service *RepoProfileService) GetFeed(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*models.Profile, error) {
	return service.repo.GetFeed(ctx, userID, limit, offset)
}
