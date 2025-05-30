package db

import (
	"fmt"

	"gorm.io/gorm"
)

var DM *gorm.DB

func Init() {
	dsn := fmt.Sprintf(
		"host=%s user=%S password=%s port=%s sslmode=disable",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "scoredb"),
		getEnv("DB_PORT", "5432"),
	)
}
