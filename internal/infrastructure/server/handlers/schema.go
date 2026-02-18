package handlers

import (
	"net/http"

	"github.com/darksuei/suei-intelligence/internal/domain/etl"
	"github.com/gin-gonic/gin"
)

func RetrieveInternalSchema(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"schema": etl.InternalSchema,
	})
	return
}