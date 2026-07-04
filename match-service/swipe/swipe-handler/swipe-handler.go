package swipe_handler

import (
	service "github.com/ZakSlinin/cofounders-match-backend/match-service/swipe/swipe-service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SwipeHandler struct {
	swipeService *service.SwipeService
}

func NewSwipeHandler(swipeService *service.SwipeService) *SwipeHandler {
	return &SwipeHandler{swipeService: swipeService}
}

func (h *SwipeHandler) Like(g *gin.Context) {
	fromUser := g.GetString("user_id")
	toUser := g.Param("user_id")

	matched, matchID, err := h.swipeService.Like(g.Request.Context(), fromUser, toUser)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"matched": matched, "match_id": matchID})
}
