package config

import (
	"github.com/darksuei/suei-intelligence/internal/domain"
)

type CacheConfig struct {
	CacheType        domain.CacheType `default:"memory"`
	RedisAddr  string             `required:"false"`
	RedisPassword    string             `required:"false"`
	RedisDB int             `required:"false"`
}
