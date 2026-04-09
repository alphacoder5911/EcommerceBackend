// Package database handles database connection setup
package database

import (
	"fmt"

	"github.com/alphacoder5911/EcommerceBackend/internal/config"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// New creates a new database connection
func New(cfg *config.DatabaseConfig) (*gorm.DB, error) {

	var dsn string

	if cfg.URL != "" {
		log.Info().Msg("url retrieved")
		dsn = cfg.URL
	} else {
		dsn = fmt.Sprintf(
			"host=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC user=%s",
			cfg.Host, cfg.Password, cfg.Name, cfg.Port, cfg.SSLMode, cfg.User,
		)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	return db, nil
}
