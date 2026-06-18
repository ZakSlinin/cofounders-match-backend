package profile_repository

import (
	"context"
	"github.com/ZakSlinin/cofounders-match-backend/user-service/models"
	"gorm.io/gorm"
)

type ProfileRepository interface {
	Create(ctx context.Context, profile *models.Profile) (*models.Profile, error)
	UpdateAvatar(ctx context.Context, userID string, avatarURL string) error
	GetByUserID(ctx context.Context, userID string) (*models.Profile, error)
}

type PostgresProfileRepository struct {
	db *gorm.DB
}

func NewPostgresProfileRepository(db *gorm.DB) *PostgresProfileRepository {
	return &PostgresProfileRepository{db: db}
}

func (repo *PostgresProfileRepository) Create(ctx context.Context, profile *models.Profile) (*models.Profile, error) {
	result := repo.db.WithContext(ctx).Create(profile)

	if result.Error != nil {
		return nil, result.Error
	}

	return profile, nil
}

func (repo *PostgresProfileRepository) UpdateAvatar(ctx context.Context, userID string, avatarURL string) error {
	result := repo.db.WithContext(ctx).
		Model(&models.Profile{}).
		Where("user_id = ?", userID).
		Update("avatar_url", avatarURL)
	return result.Error
}

func (repo *PostgresProfileRepository) GetByUserID(ctx context.Context, userID string) (*models.Profile, error) {
	result := repo.db.WithContext(ctx).Model(&models.Profile{}).Where("user_id = ?", userID)

	if result.Error != nil {
		return nil, result.Error
	}

	return &models.Profile{}, result.Error
}
