package database

import (
	"github.com/darksuei/suei-intelligence/internal/config"
	"github.com/darksuei/suei-intelligence/internal/domain"
	"github.com/darksuei/suei-intelligence/internal/infrastructure/database/postgres"
	"github.com/darksuei/suei-intelligence/internal/infrastructure/database/sqlite"
)

func InitializeConnection(config *config.DatabaseConfig) {
	switch config.DatabaseType {
		case domain.DatabaseTypePostgres:
			postgres.Connect(config)
		case domain.DatabaseTypeSqlite:
			sqlite.Connect(config)
		default:
			// Use SQLite as Default
			sqlite.Connect(config)
	}
}
