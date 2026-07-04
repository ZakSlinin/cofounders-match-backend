package match_repository

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type Match struct {
	ID        string `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	User1ID   string `gorm:"type:uuid;not null"`
	User2ID   string `gorm:"type:uuid;not null"`
	CreatedAt time.Time
}

type MatchRepository struct {
	db *gorm.DB
}

func NewMatchRepository(db *gorm.DB) *MatchRepository {
	return &MatchRepository{db: db}
}

func (r *MatchRepository) CreateMatch(ctx context.Context, user1, user2 string) (string, error) {
	match := &Match{User1ID: user1, User2ID: user2}
	err := r.db.WithContext(ctx).Create(match).Error
	return match.ID, err
}

func (r *MatchRepository) GetMatchesByUser(ctx context.Context, userID string) ([]Match, error) {
	var matches []Match
	err := r.db.WithContext(ctx).
		Where("user1_id = ? OR user2_id = ?", userID, userID).
		Find(&matches).Error
	return matches, err
}

func (r *MatchRepository) DeleteMatch(ctx context.Context, matchID, userID string) error {
	return r.db.WithContext(ctx).
		Where("id = ? AND (user1_id = ? OR user2_id = ?)", matchID, userID, userID).
		Delete(&Match{}).Error
}
