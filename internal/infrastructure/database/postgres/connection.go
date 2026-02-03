package postgres

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/darksuei/suei-intelligence/internal/config"
)

var DB *gorm.DB

func ValidateConfig(c *config.DatabaseConfig) error {
	if c.DatabaseHost == "" {
		return errors.New("DATABASE_HOST is required")
	}
	if c.DatabasePort == "" {
		return errors.New("DATABASE_PORT is required")
	}
	if c.DatabaseUsername == "" {
		return errors.New("DATABASE_USERNAME is required")
	}
	if c.DatabasePassword == "" {
		return errors.New("DATABASE_PASSWORD is required")
	}
	if c.DatabaseName == "" {
		return errors.New("DATABASE_NAME is required")
	}
	return nil
}

func Connect(config *config.DatabaseConfig) {
	if err := ValidateConfig(config); err != nil {
		log.Fatalf("Invalid postgres config: %v", err)
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", 
		config.DatabaseHost, config.DatabasePort, 
		config.DatabaseUsername, config.DatabasePassword, 
		config.DatabaseName, strconv.FormatBool(config.DatabaseUseSSL))
	
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to postgres: %s", err)
	}

	log.Printf("Successfully connected to postgres..")
}
