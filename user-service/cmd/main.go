package main

import (
	"fmt"
	auth_handler "github.com/ZakSlinin/cofounders-match-backend/user-service/auth/auth-handler"
	auth_service "github.com/ZakSlinin/cofounders-match-backend/user-service/auth/auth-service"
	"github.com/ZakSlinin/cofounders-match-backend/user-service/cmd/config"
	db "github.com/ZakSlinin/cofounders-match-backend/user-service/cmd/db"
	user_repository "github.com/ZakSlinin/cofounders-match-backend/user-service/user/user-repository"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config.Load()

	gormDB, err := db.NewPostgres(cfg)
	if err != nil {
		log.Fatal(err)
	}

	authRepo := user_repository.NewPostgresUserRepository(gormDB)
	authService := auth_service.NewAuthService(authRepo)
	authHandler := auth_handler.NewAuthHandler(authService)

	r := gin.Default()

	auth := r.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		//TODO: auth.POST("/refresh", authHandler.Refresh)
	}

	port := cfg.Port
	fmt.Println("Starting server on port", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
