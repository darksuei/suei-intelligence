package handlers

import (
	"log"
	"net/http"
	"strconv"

	authorizationService "github.com/darksuei/suei-intelligence/internal/application/authorization"
	"github.com/darksuei/suei-intelligence/internal/application/datasource"
	"github.com/darksuei/suei-intelligence/internal/config"
	authorizationDomain "github.com/darksuei/suei-intelligence/internal/domain/authorization"
	"github.com/darksuei/suei-intelligence/internal/infrastructure/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
)

func NewDatasource(c *gin.Context) {
	var req struct {
		// datasource name/identifier
		// datasource connection details
	}

	key := c.Param("key") // assumes route is like /projects/:key
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Project key is required",
		})
		return
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "Invalid request: Missing required fields.",
		})
		return
	}

	var databaseCfg config.DatabaseConfig
	if err := envconfig.Process("", &databaseCfg); err != nil {
		log.Fatalf("Failed to load database config: %v", err)
	}

	createdByEmail, err := utils.GetUserEmailFromContext(c)

	if err != nil || createdByEmail == nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get account",
		})
		return
	}

	// Authorization
	allow, err := authorizationService.EnforceRoles(utils.GetUserRolesFromContext(c), "org", authorizationDomain.Organization, "write")

	if err != nil || !allow {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "forbidden",
		})
		return
	}

	// Create datasource
	_datasource, err := datasource.NewDatasource(key, *createdByEmail, &databaseCfg)

	if err != nil {
		log.Printf("Error creating datasource: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"datasource": _datasource,
		})
}

func RetrieveDatasources(c *gin.Context) {
	key := c.Param("key") // assumes route is like /projects/:key
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Project key is required",
		})
		return
	}

	var databaseCfg config.DatabaseConfig
	if err := envconfig.Process("", &databaseCfg); err != nil {
		log.Fatalf("Failed to load database config: %v", err)
	}

	// Authorization
	allow, err := authorizationService.EnforceRoles(utils.GetUserRolesFromContext(c), "org", authorizationDomain.Organization, "read")

	if err != nil || !allow {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "forbidden",
		})
		return
	}

	// Retrieve datasources
	_datasources, err := datasource.RetrieveDatasources(key, &databaseCfg)

	if err != nil {
		log.Printf("Error retrieving datasources: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"datasources": _datasources,
	})
}

func DeleteDatasource (c *gin.Context) {
	key := c.Param("key") // assumes route is like /projects/:key
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Project key is required",
		})
		return
	}

	idParam := c.Param("id") // /projects/:id
	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Project id is required",
		})
		return
	}

	datasourceID, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid datasource id",
		})
		return
	}

	var databaseCfg config.DatabaseConfig
	if err := envconfig.Process("", &databaseCfg); err != nil {
		log.Fatalf("Failed to load database config: %v", err)
	}

	// Authorization
	allow, err := authorizationService.EnforceRoles(utils.GetUserRolesFromContext(c), "org", authorizationDomain.Organization, "write")

	if err != nil || !allow {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "forbidden",
		})
		return
	}

	// Delete datasource
	err = datasource.SoftDeleteDatasource(uint(datasourceID), key, &databaseCfg)

	if err != nil {
		log.Printf("Error deleting datasource: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}