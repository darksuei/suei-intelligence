package handlers

import (
	"net/http"

	"github.com/darksuei/suei-intelligence/internal/domain/schema"
	"github.com/gin-gonic/gin"
)

func RetrieveInternalSchema(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"schema": schema.InternalSchema,
	})
	return
}