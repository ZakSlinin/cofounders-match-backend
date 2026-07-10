package repository

import (
	"context"
	swipe "github.com/ZakSlinin/cofounders-match-backend/match-service/models"
	"gorm.io/gorm"
)

type SwipeRepository struct {
	db *gorm.DB
}

func NewSwipeRepository(db *gorm.DB) *SwipeRepository {
	return &SwipeRepository{db: db}
}

func (r *SwipeRepository) CreateSwipe(ctx context.Context, fromUser, toUser string) error {
	return r.db.WithContext(ctx).Create(&swipe.Swipe{
		FromUser: fromUser,
		ToUser:   toUser,
	}).Error
}

func (r *SwipeRepository) CheckMutualLike(ctx context.Context, fromUser, toUser string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&swipe.Swipe{}).
		Where("from_user = ? AND to_user = ?", toUser, fromUser).
		Count(&count).Error
	return count > 0, err
}
