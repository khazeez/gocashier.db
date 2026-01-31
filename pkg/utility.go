package pkg

import (
	"log"
	"os"

	"gocashier.db/internal/models"

	"github.com/joho/godotenv"
)

var cfg *models.Config

func Load() *models.Config {
	if cfg != nil {
		return cfg
	}

	if err := godotenv.Load(); err != nil {
		log.Println(".env not found, using system environment")
	}

	cfg = &models.Config{
		AppPort: getEnv("APP_PORT"),
		DBURL:   getEnv("DATABASE_URL"),
	}

	return cfg
}

func getEnv(key string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return ""
}
