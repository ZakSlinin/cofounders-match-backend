package match_handler

import (
	service "github.com/ZakSlinin/cofounders-match-backend/match-service/match/match-service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MatchHandler struct {
	matchService *service.MatchService
}

func NewMatchHandler(matchService *service.MatchService) *MatchHandler {
	return &MatchHandler{matchService: matchService}
}

func (h *MatchHandler) GetMatches(g *gin.Context) {
	userID := g.GetString("user_id")

	matches, err := h.matchService.GetMatches(g.Request.Context(), userID)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, matches)
}

func (h *MatchHandler) DeleteMatch(g *gin.Context) {
	userID := g.GetString("user_id")
	matchID := g.Param("id")

	if err := h.matchService.DeleteMatch(g.Request.Context(), matchID, userID); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"deleted": true})
}
