package main

import (
	"fmt"
	auth_handler "github.com/ZakSlinin/cofounders-match-backend/user-service/auth/auth-handler"
	auth_service "github.com/ZakSlinin/cofounders-match-backend/user-service/auth/auth-service"
	"github.com/ZakSlinin/cofounders-match-backend/user-service/cmd/config"
	db "github.com/ZakSlinin/cofounders-match-backend/user-service/cmd/db"
	profile_handler "github.com/ZakSlinin/cofounders-match-backend/user-service/profile/profile-handler"
	profile_repository "github.com/ZakSlinin/cofounders-match-backend/user-service/profile/profile-repository"
	profile_service "github.com/ZakSlinin/cofounders-match-backend/user-service/profile/profile-service"
	"github.com/ZakSlinin/cofounders-match-backend/user-service/profile/storage"
	vision_service "github.com/ZakSlinin/cofounders-match-backend/user-service/profile/vision-service"
	user_repository "github.com/ZakSlinin/cofounders-match-backend/user-service/user/user-repository"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func runMigrations(dsn string) error {
	m, err := migrate.New("file://migrations", dsn)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file, reading from environment")
	}

	cfg := config.Load()

	if err := runMigrations(cfg.DBURL); err != nil {
		log.Fatal("migrations failed:", err)
	}

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
	visionService := vision_service.NewVisionService(os.Getenv("VISION_API_KEY"), os.Getenv("VISION_FOLDER_ID"))
	profileHanlder := profile_handler.NewProfileHandler(profileService, storageService, visionService)

	r := gin.Default()

	//TODO: only gateway instead of all domains
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
	{
		protected.POST("/profiles", profileHanlder.CreateProfile)
		protected.POST("/profiles/avatar", profileHanlder.UploadAvatar)
		protected.GET("/profiles/me", profileHanlder.GetMe)
		protected.PATCH("/profiles/me", profileHanlder.UpdateProfile)

		protected.GET("/feed", profileHanlder.GetFeed)
	}

	r.GET("/profiles/:user_id", profileHanlder.GetByUserID)

	port := cfg.Port
	fmt.Println("Starting server on port", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
