package profile_repository

import (
	"context"
	"github.com/ZakSlinin/cofounders-match-backend/user-service/models"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
	"log"
)

type ProfileRepository interface {
	Create(ctx context.Context, profile *models.Profile) (*models.Profile, error)
	UpdateAvatar(ctx context.Context, userID string, avatarURL string) error
	GetByUserID(ctx context.Context, userID uuid.UUID) (*models.Profile, error)
	UpdateProfile(ctx context.Context, userID uuid.UUID, update *models.UpdateProfileRequest) (*models.Profile, error)
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

func (repo *PostgresProfileRepository) GetByUserID(ctx context.Context, userID uuid.UUID) (*models.Profile, error) {
	var profile models.Profile
	result := repo.db.WithContext(ctx).Where("user_id = ?", userID).Take(&profile)

	log.Printf("GetByUserID: userID=%s, error=%v, profile=%+v", userID, result.Error, profile)

	if result.Error != nil {
		return nil, result.Error
	}

	return &profile, nil
}

func (repo *PostgresProfileRepository) UpdateProfile(ctx context.Context, userID uuid.UUID, update *models.UpdateProfileRequest) (*models.Profile, error) {
	profile, err := repo.GetByUserID(ctx, userID)
	if err != nil || profile == nil {
		return nil, err
	}

	if update.Name != nil {
		profile.Name = *update.Name
	}
	if update.Bio != nil {
		profile.Bio = *update.Bio
	}
	if update.City != nil {
		profile.City = *update.City
	}
	if update.Skills != nil {
		profile.Skills = pq.StringArray(update.Skills)
	}
	if update.LookingFor != nil {
		profile.LookingFor = pq.StringArray(update.LookingFor)
	}
	if update.AvailableForProjects != nil {
		profile.AvailableForProjects = *update.AvailableForProjects
	}

	repo.db.WithContext(ctx).Save(profile)
	return profile, nil
}
