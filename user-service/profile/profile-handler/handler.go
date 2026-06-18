package profile_handler

import (
	"github.com/ZakSlinin/cofounders-match-backend/user-service/models"
	profile_service "github.com/ZakSlinin/cofounders-match-backend/user-service/profile/profile-service"
	storage "github.com/ZakSlinin/cofounders-match-backend/user-service/profile/storage"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type ProfileHandler struct {
	storage *storage.StorageService
	service profile_service.ProfileService
}

func NewProfileHandler(service profile_service.ProfileService, storage *storage.StorageService) *ProfileHandler {
	return &ProfileHandler{service: service, storage: storage}
}

func (h *ProfileHandler) CreateProfile(g *gin.Context) {
	var req models.CreateProfileRequest

	userIDstr := g.GetString("user_id")
	userID, err := uuid.Parse(userIDstr)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	profile, createdErr := h.service.Create(g.Request.Context(), &models.Profile{
		ID:                   uuid.New(),
		UserID:               userID,
		Name:                 req.Name,
		Bio:                  req.Bio,
		AvatarURL:            req.AvatarURL,
		City:                 req.City,
		LookingFor:           req.LookingFor,
		Skills:               req.Skills,
		AvailableForProjects: req.AvailableForProjects,
		CreatedAt:            time.Now(),
	})

	if createdErr != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": createdErr.Error()})
		return
	}

	g.JSON(http.StatusCreated, gin.H{"profile": profile})
}

func (h *ProfileHandler) UploadAvatar(g *gin.Context) {
	userID := g.GetString("user_id")

	file, header, err := g.Request.FormFile("file")
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "no file"})
		return
	}
	defer file.Close()

	url, err := h.storage.Upload(g.Request.Context(), file, header)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.service.UpdateAvatar(g.Request.Context(), userID, url)
	g.JSON(http.StatusOK, gin.H{"avatar_url": url})
}
