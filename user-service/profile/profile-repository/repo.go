package profile_repository

import (
	"context"
	"github.com/ZakSlinin/cofounders-match-backend/user-service/models"
	"gorm.io/gorm"
)

type ProfileRepository interface {
	Create(ctx context.Context, user *models.Profile) (*models.Profile, error)
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
