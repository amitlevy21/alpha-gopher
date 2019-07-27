package api

import (
	"net/http"

	"github.com/amitlevy21/alpha-gopher/internal"
	"github.com/gin-gonic/gin"
)

func Find(c *gin.Context) {
	path := c.Query("path")
	opts := c.QueryArray("opts")

	combinedStd, err := internal.Find(path, opts)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "std": combinedStd})
		return
	}
	c.JSON(http.StatusOK, gin.H{"matches": combinedStd})
}
