package api

import (
	"net/http"
	"time"

	"github.com/amitlevy21/alpha-gopher/internal"
	"github.com/gin-gonic/gin"
)

func GetDate(c *gin.Context) {
	date, err := internal.GetDate()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"date": date})
}

func SetDate(c *gin.Context) {
	newDate := c.Param("date")
	t, err := time.Parse("02-Jan-2006", newDate)
	combinedStd, err := internal.SetDate(t)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "stderr": combinedStd})
		return
	}
	c.JSON(http.StatusOK, gin.H{"date": combinedStd})
}
