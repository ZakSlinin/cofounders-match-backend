package auth_service

import (
	"context"
	"errors"
	"github.com/ZakSlinin/cofounders-match-backend/user-service/models"
	user_repository "github.com/ZakSlinin/cofounders-match-backend/user-service/user/user-repository"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type AuthService interface {
	Register(ctx context.Context, email, password, role string) (*models.User, string, string, error)
	Login(ctx context.Context, email, password string) (*models.User, string, string, error)
}

type authService struct {
	repo user_repository.UserRepository
}

func NewAuthService(repo user_repository.UserRepository) *authService {
	return &authService{repo: repo}
}

func (s *authService) Register(ctx context.Context, email, password, role string) (*models.User, string, string, error) {
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, "", "", err
	}

	if user != nil {
		return nil, "", "", errors.New("email already exists")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 4)

	user = &models.User{
		ID:           uuid.New(),
		Email:        email,
		PasswordHash: string(passwordHash),
		Role:         role,
	}

	user, err = s.repo.Create(ctx, user)
	if err != nil {
		return nil, "", "", err
	}

	accessToken, refreshToken, err := generateJWT(user)
	if err != nil {
		return nil, "", "", err
	}

	err = s.repo.SaveTokens(ctx, user.ID, refreshToken)

	return user, accessToken, refreshToken, nil
}

func (s *authService) Login(ctx context.Context, email, password string) (*models.User, string, string, error) {
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return user, "", "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, "", "", errors.New("invalid credentials")
	}

	accessToken, refreshToken, err := generateJWT(user)

	err = s.repo.SaveTokens(ctx, user.ID, refreshToken)

	return user, accessToken, refreshToken, nil
}

func generateJWT(user *models.User) (access, refresh string, err error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	access, err = accessToken.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", "", err
	}

	refreshClaims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(),
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refresh, err = refreshToken.SignedString([]byte(os.Getenv("JWT_REFRESH_SECRET")))

	return
}
