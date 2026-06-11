package auth_handler

import (
	auth_service "github.com/ZakSlinin/cofounders-match-backend/user-service/auth/auth-service"
	"github.com/ZakSlinin/cofounders-match-backend/user-service/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type authHandler struct {
	authService auth_service.AuthService
}

func NewAuthHandler(authService auth_service.AuthService) *authHandler {
	return &authHandler{authService: authService}
}

func (h *authHandler) Register(g *gin.Context) {
	var req models.RegisterRequest

	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, accessToken, refreshToken, registerErr := h.authService.Register(g.Request.Context(), req.Email, req.Password, req.Role)

	if registerErr != nil {
		if registerErr.Error() == "email already exists" {
			g.JSON(http.StatusBadRequest, gin.H{"error": registerErr.Error()})
			return
		}
		g.JSON(http.StatusInternalServerError, gin.H{"error": registerErr.Error()})
		return
	}

	g.JSON(http.StatusOK, models.AuthResponse{ID: createdUser.ID, Email: createdUser.Email, Role: createdUser.Role, RefreshToken: refreshToken, AccessToken: accessToken})
}
