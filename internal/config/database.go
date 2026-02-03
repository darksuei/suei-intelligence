package config

import (
	"github.com/darksuei/suei-intelligence/internal/domain"
)

type DatabaseConfig struct {
	DatabaseType        domain.DatabaseType `default:"sqlite"`
	DatabaseHost  string             `required:"false"`
	DatabasePort    string             `required:"false"`
	DatabaseUsername string             `required:"false"`
	DatabasePassword string             `required:"false"`
	DatabaseUseSSL    bool               `default:"false"`
	DatabaseName    string             `required:"false"`
	DatabasePath    string             `default:"./data/app.db"`
}
