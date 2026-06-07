package user_repository

import (
	"context"
	"github.com/ZakSlinin/cofounders-match-backend/user-service/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
}

type PostgresUserRepository struct {
	db *gorm.DB
}

func (repo *PostgresUserRepository) Create(ctx context.Context, user *models.User) (*models.User, error) {
	result := repo.db.WithContext(ctx).Create(user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (repo *PostgresUserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	result := repo.db.WithContext(ctx).Where("email = ?", email).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
