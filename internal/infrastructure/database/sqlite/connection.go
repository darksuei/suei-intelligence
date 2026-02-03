package sqlite

import (
	"errors"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/darksuei/suei-intelligence/internal/config"
)

var DB *gorm.DB

func ValidateConfig(c *config.DatabaseConfig) error {
	if c.DatabasePath == "" {
		return errors.New("DATABASE_PATH is required")
	}
	return nil
}

func Connect(cfg *config.DatabaseConfig) {
	if err := ValidateConfig(cfg); err != nil {
		log.Fatalf("Invalid sqlite config: %v", err)
	}

	var err error

	DB, err = gorm.Open(sqlite.Open(cfg.DatabasePath), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to sqlite: %v", err)
	}

	log.Println("Successfully connected to sqlite")
}
