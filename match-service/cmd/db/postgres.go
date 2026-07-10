package db

import (
	"github.com/ZakSlinin/cofounders-match-backend/match-service/cmd/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres(cfg *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.DBURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
