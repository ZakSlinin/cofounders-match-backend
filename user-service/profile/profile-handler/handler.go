package profile_handler

import (
	"fmt"
	"github.com/ZakSlinin/cofounders-match-backend/user-service/models"
	profile_service "github.com/ZakSlinin/cofounders-match-backend/user-service/profile/profile-service"
	storage "github.com/ZakSlinin/cofounders-match-backend/user-service/profile/storage"
	vision_service "github.com/ZakSlinin/cofounders-match-backend/user-service/profile/vision-service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"net/http"
	"time"
)

type ProfileHandler struct {
	storage       *storage.StorageService
	service       profile_service.ProfileService
	visionService vision_service.VisionService
}

func NewProfileHandler(service profile_service.ProfileService, storage *storage.StorageService, visionService vision_service.VisionService) *ProfileHandler {
	return &ProfileHandler{service: service, storage: storage, visionService: visionService}
}

func (h *ProfileHandler) CreateProfile(g *gin.Context) {
	var req models.CreateProfileRequest

	userIDstr := g.GetHeader("X-User-ID")
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
		City:                 req.City,
		LookingFor:           pq.StringArray(req.LookingFor),
		Skills:               pq.StringArray(req.Skills),
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
	userID := g.GetHeader("X-User-ID")

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

	safe, err := h.visionService.CheckImage(g.Request.Context(), url)
	if !safe {
		h.storage.Delete(g.Request.Context(), url) // удаляем из S3
		g.JSON(http.StatusBadRequest, gin.H{"error": "inappropriate content"})
		return
	}

	h.service.UpdateAvatar(g.Request.Context(), userID, url)
	g.JSON(http.StatusOK, gin.H{"avatar_url": url})
}

func (h *ProfileHandler) GetByUserID(g *gin.Context) {
	userIDstr := g.Param("user_id")
	userID, err := uuid.Parse(userIDstr)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	profile, err := h.service.GetByUserID(g.Request.Context(), userID)

	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if profile == nil {
		g.JSON(http.StatusNotFound, gin.H{"error": "profile not found"})
		return
	}

	g.JSON(http.StatusOK, gin.H{"profile": profile})
}

func (h *ProfileHandler) GetMe(g *gin.Context) {
	userID, err := uuid.Parse(g.GetHeader("X-User-ID"))
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	profile, err := h.service.GetByUserID(g.Request.Context(), userID)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if profile == nil {
		g.JSON(http.StatusNotFound, gin.H{"error": "profile not found"})
		return
	}

	g.JSON(http.StatusOK, gin.H{"profile": profile})
}

func (h *ProfileHandler) UpdateProfile(g *gin.Context) {
	userID, err := uuid.Parse(g.GetHeader("X-User-ID"))

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req models.UpdateProfileRequest

	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	profile, updatedProfileErr := h.service.UpdateProfile(g.Request.Context(), userID, &req)

	if updatedProfileErr != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": updatedProfileErr.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"profile": profile})
}

func (h *ProfileHandler) GetFeed(g *gin.Context) {
	userID, err := uuid.Parse(g.GetHeader("X-User-ID"))
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	limit := 20
	offset := 0

	if l := g.Query("limit"); l != "" {
		fmt.Sscan(l, &limit)
	}
	if o := g.Query("offset"); o != "" {
		fmt.Sscan(o, &offset)
	}

	profiles, err := h.service.GetFeed(g.Request.Context(), userID, limit, offset)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"profiles": profiles, "limit": limit, "offset": offset})
}
