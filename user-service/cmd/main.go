package main

import (
	"fmt"
	auth_handler "github.com/ZakSlinin/cofounders-match-backend/user-service/auth/auth-handler"
	auth_service "github.com/ZakSlinin/cofounders-match-backend/user-service/auth/auth-service"
	"github.com/ZakSlinin/cofounders-match-backend/user-service/cmd/config"
	db "github.com/ZakSlinin/cofounders-match-backend/user-service/cmd/db"
	"github.com/ZakSlinin/cofounders-match-backend/user-service/cmd/middleware"
	profile_handler "github.com/ZakSlinin/cofounders-match-backend/user-service/profile/profile-handler"
	profile_repository "github.com/ZakSlinin/cofounders-match-backend/user-service/profile/profile-repository"
	profile_service "github.com/ZakSlinin/cofounders-match-backend/user-service/profile/profile-service"
	"github.com/ZakSlinin/cofounders-match-backend/user-service/profile/storage"
	user_repository "github.com/ZakSlinin/cofounders-match-backend/user-service/user/user-repository"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file, reading from environment")
	}

	cfg := config.Load()

	gormDB, err := db.NewPostgres(cfg)
	if err != nil {
		log.Fatal(err)
	}

	authRepo := user_repository.NewPostgresUserRepository(gormDB)
	authService := auth_service.NewAuthService(authRepo)
	authHandler := auth_handler.NewAuthHandler(authService)

	// --- profile ---
	profileRepo := profile_repository.NewPostgresProfileRepository(gormDB)
	profileService := profile_service.NewProfileService(profileRepo)

	// --- s3 client ---
	s3Client := db.NewS3Client()

	// --- storage service ---
	storageService := storage.NewStorageService(s3Client, os.Getenv("YC_BUCKET"))
	profileHanlder := profile_handler.NewProfileHandler(profileService, storageService)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	auth := r.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		auth.POST("/refresh", authHandler.Refresh)
	}

	protected := r.Group("/")
	protected.Use(middleware.JWTMiddleware())
	{
		protected.POST("/profiles", profileHanlder.CreateProfile)
		protected.POST("/profiles/avatar", profileHanlder.UploadAvatar)
	}

	port := cfg.Port
	fmt.Println("Starting server on port", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
