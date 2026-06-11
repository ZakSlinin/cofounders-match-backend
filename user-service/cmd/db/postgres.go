package db

import (
	"github.com/ZakSlinin/cofounders-match-backend/user-service/cmd/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres(ccfg *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(ccfg.DBURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
