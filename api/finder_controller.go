package api

import (
	"net/http"

	"github.com/amitlevy21/alpha-gopher/internal"
	"github.com/gin-gonic/gin"
)

func Find(c *gin.Context) {
	path := c.Query("path")
	opts := c.QueryArray("opts")

	combinedStd, _ := internal.Find(path, opts)
	c.JSON(http.StatusOK, gin.H{"matches": combinedStd})
}
