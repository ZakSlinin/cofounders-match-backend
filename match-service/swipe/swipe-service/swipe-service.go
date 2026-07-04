package swipe_service

import (
	"context"
	match_repository "github.com/ZakSlinin/cofounders-match-backend/match-service/match/match-repository"
	swipe_repository "github.com/ZakSlinin/cofounders-match-backend/match-service/swipe/swipe-repository"
)

type SwipeService struct {
	swipeRepo *swipe_repository.SwipeRepository
	matchRepo *match_repository.MatchRepository
}

func NewSwipeService(swipeRepo *swipe_repository.SwipeRepository, matchRepo *match_repository.MatchRepository) *SwipeService {
	return &SwipeService{swipeRepo: swipeRepo, matchRepo: matchRepo}
}

func (s *SwipeService) Like(ctx context.Context, fromUser, toUser string) (bool, string, error) {
	if err := s.swipeRepo.CreateSwipe(ctx, fromUser, toUser); err != nil {
		return false, "", err
	}

	mutual, err := s.swipeRepo.CheckMutualLike(ctx, fromUser, toUser)
	if err != nil {
		return false, "", err
	}

	if mutual {
		matchID, err := s.matchRepo.CreateMatch(ctx, fromUser, toUser)
		if err != nil {
			return false, "", err
		}
		return true, matchID, nil
	}

	return false, "", nil
}
