package user_service

import (
	"context"
	"github.com/ZakSlinin/cofounders-match-backend/user-service/models"
	user_repository "github.com/ZakSlinin/cofounders-match-backend/user-service/user/user-repository"
)

type UserService struct {
	repo user_repository.UserRepository
}

func NewUserService(repo *user_repository.UserRepository) *UserService {
	return &UserService{repo: *repo}
}

func (service *UserService) Create(ctx context.Context, user *models.User) (*models.User, error) {
	return service.repo.Create(ctx, user)
}

func (service *UserService) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	return service.repo.GetByEmail(ctx, email)
}
