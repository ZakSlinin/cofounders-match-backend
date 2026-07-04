package match_service

import (
	"context"
	repository "github.com/ZakSlinin/cofounders-match-backend/match-service/match/match-repository"
)

type MatchService struct {
	matchRepo *repository.MatchRepository
}

func NewMatchService(matchRepo *repository.MatchRepository) *MatchService {
	return &MatchService{matchRepo: matchRepo}
}

func (s *MatchService) GetMatches(ctx context.Context, userID string) ([]repository.Match, error) {
	return s.matchRepo.GetMatchesByUser(ctx, userID)
}

func (s *MatchService) DeleteMatch(ctx context.Context, matchID, userID string) error {
	return s.matchRepo.DeleteMatch(ctx, matchID, userID)
}
