package api

import (
	"net/http"

	"github.com/amitlevy21/alpha-gopher/internal"
	"github.com/gin-gonic/gin"
)

func GetDevices(c *gin.Context) {
	devices, err := internal.GetDevices()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"devices": devices})
}

func Backup(c *gin.Context) {
	path := c.Query("path")
	_, err := internal.Backup(path)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
