package handlers

import (
	"log"
	"net/http"

	"github.com/darksuei/suei-intelligence/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
)

func RetrieveConfig(c *gin.Context) {
	var commonCfg config.CommonConfig
	if err := envconfig.Process("", &commonCfg); err != nil {
		log.Fatalf("Failed to load common config: %v", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"enforce_mfa": commonCfg.EnforceMfa,
	})
	return
}