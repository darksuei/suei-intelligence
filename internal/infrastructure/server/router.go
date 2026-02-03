package server

import (
	"github.com/darksuei/suei-intelligence/internal/infrastructure/server/handlers"
	"github.com/gin-gonic/gin"
)

func InitializeRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/health", handlers.Health)

	return router
}
