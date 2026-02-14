package handlers

import (
	"log"
	"net/http"

	"github.com/darksuei/suei-intelligence/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
)

func Health(c *gin.Context) {
	log.Print("Received health check request..")

	var commonCfg config.CommonConfig
	err := envconfig.Process("", &commonCfg)
	if err != nil {
		log.Fatalf("Failed to load common config: %v", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Healthy",
		"version": "v0.0.0",
		"copyright": "2026 su3i inc.",
		"environment": commonCfg.AppEnv,
	  })
}