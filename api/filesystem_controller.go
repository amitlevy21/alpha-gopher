package api

import (
	"net/http"

	"github.com/amitlevy21/alpha-gopher/internal"
	"github.com/gin-gonic/gin"
)

func Copy(c *gin.Context) {
	src := c.Query("src")
	dst := c.Query("dst")
	opts := c.QueryArray("opts")

	combinedStd, err := internal.Copy(src, dst, opts)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "std": combinedStd})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func Move(c *gin.Context) {
	src := c.Query("src")
	dst := c.Query("dst")

	combinedStd, err := internal.Move(src, dst)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "std": combinedStd})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func Remove(c *gin.Context) {
	dst := c.Query("dst")

	combinedStd, err := internal.Remove(dst)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "std": combinedStd})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func List(c *gin.Context) {
	path := c.Query("path")
	opts := c.QueryArray("opts")

	combinedStd, err := internal.List(path, opts)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "std": combinedStd})
		return
	}
	c.JSON(http.StatusOK, gin.H{"entries": combinedStd})
}

func ListAll(c *gin.Context) {
	filesystem, err := internal.BuildTree()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "std": filesystem})
		return
	}
	c.JSON(http.StatusOK, gin.H{"std": filesystem})
}
