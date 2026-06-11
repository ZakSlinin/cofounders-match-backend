package user_repository

import (
	"context"
	"errors"
	"github.com/ZakSlinin/cofounders-match-backend/user-service/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	SaveTokens(ctx context.Context, userID uuid.UUID, token string) error
	GetRefreshToken(ctx context.Context, token string) (*models.RefreshToken, error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.User, error)
}

type PostgresUserRepository struct {
	db *gorm.DB
}

func NewPostgresUserRepository(db *gorm.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
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
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &user, nil
}

func (repo *PostgresUserRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	var user models.User
	result := repo.db.WithContext(ctx).Where("id = ?", id).Take(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &user, nil
}

func (repo *PostgresUserRepository) SaveTokens(ctx context.Context, userID uuid.UUID, token string) error {
	refreshToken := &models.RefreshToken{
		UserID:    userID,
		Token:     token,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 30),
	}

	result := repo.db.WithContext(ctx).Create(refreshToken)
	return result.Error
}

func (repo *PostgresUserRepository) GetRefreshToken(ctx context.Context, token string) (*models.RefreshToken, error) {
	var refreshToken models.RefreshToken
	result := repo.db.WithContext(ctx).Where("token = ?", token).Take(&refreshToken)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &refreshToken, nil
}
