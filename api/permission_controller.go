package api

import (
	"net/http"

	"github.com/amitlevy21/alpha-gopher/internal"

	"github.com/gin-gonic/gin"
)

func GetOwner(c *gin.Context) {
	path := c.Query("path")
	combinedStd, err := internal.GetOwner(path)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "stderr": combinedStd})
		return
	}
	c.JSON(http.StatusOK, gin.H{"path": path, "owner": combinedStd})
}

func ChangeOwnership(c *gin.Context) {
	owner := c.Param("newOwner")
	path := c.Query("path")
	err := internal.ChangeOwnership(path, owner)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func GetCurrentPermissions(c *gin.Context) {
	path := c.Query("path")
	pathMode, err := internal.GetCurrentPermissions(path)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"path": path, "mode": pathMode})
}
