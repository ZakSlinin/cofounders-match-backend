package config

import "os"

type Config struct {
	DBUser         string
	DBPass         string
	DBHost         string
	DBPort         string
	DBName         string
	DBURL          string
	Port           string
	JWTSecret      string
	VisionAPIKey   string
	VisionFolderID string
}

func Load() *Config {
	cfg := &Config{
		DBUser:         os.Getenv("DB_USER"),
		DBPass:         os.Getenv("DB_PASS"),
		DBHost:         os.Getenv("DB_HOST"),
		DBPort:         os.Getenv("DB_PORT"),
		DBName:         os.Getenv("DB_NAME"),
		DBURL:          os.Getenv("DB_URL"),
		Port:           os.Getenv("PORT"),
		JWTSecret:      os.Getenv("JWT_SECRET"),
		VisionAPIKey:   os.Getenv("VISION_API_KEY"),
		VisionFolderID: os.Getenv("VISION_FOLDER_ID"),
	}

	cfg.DBURL = "postgres://" + cfg.DBUser + ":" + cfg.DBPass +
		"@" + cfg.DBHost + ":" + cfg.DBPort + "/" + cfg.DBName + "?sslmode=disable"

	return cfg
}
