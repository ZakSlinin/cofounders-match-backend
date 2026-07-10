package main

import (
	"github.com/ZakSlinin/cofounders-match-backend/match-service/cmd/config"
	db "github.com/ZakSlinin/cofounders-match-backend/match-service/cmd/db"
	match_handler "github.com/ZakSlinin/cofounders-match-backend/match-service/match/match-handler"
	match_repository "github.com/ZakSlinin/cofounders-match-backend/match-service/match/match-repository"
	match_service "github.com/ZakSlinin/cofounders-match-backend/match-service/match/match-service"
	swipe_handler "github.com/ZakSlinin/cofounders-match-backend/match-service/swipe/swipe-handler"
	swipe_repository "github.com/ZakSlinin/cofounders-match-backend/match-service/swipe/swipe-repository"
	swipe_service "github.com/ZakSlinin/cofounders-match-backend/match-service/swipe/swipe-service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found")
	}

	cfg := config.Load()

	gormDB, err := db.NewPostgres(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// --- match ---
	matchRepo := match_repository.NewMatchRepository(gormDB)
	matchService := match_service.NewMatchService(matchRepo)
	matchHandler := match_handler.NewMatchHandler(matchService)

	// --- swipe ---
	swipeRepo := swipe_repository.NewSwipeRepository(gormDB)
	swipeService := swipe_service.NewSwipeService(swipeRepo, matchRepo)
	swipeHandler := swipe_handler.NewSwipeHandler(swipeService)

	r := gin.Default()

	// --- cors ---
	//TODO: only gateway instead of all domains
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	swipe := r.Group("/swipe")
	{
		swipe.POST("/like/:user_id", swipeHandler.Like)
	}

	match := r.Group("/match")
	{
		match.GET("/matches", matchHandler.GetMatches)
		match.DELETE("/matches/:match_id", matchHandler.DeleteMatch)
	}
}
