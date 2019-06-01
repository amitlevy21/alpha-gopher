package api

import (
	"fmt"
	"net/http"

	"github.com/amitlevy21/alpha-gopher/internal"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	users, err := internal.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func AddUser(c *gin.Context) {
	userName := c.Param("name")
	err := internal.AddUser(userName)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": fmt.Sprintf("user %s created", userName)})
}

func DeleteUser(c *gin.Context) {
	userName := c.Param("name")
	err := internal.DeleteUser(userName)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": fmt.Sprintf("user %s deleted", userName)})
}

func ChangeUserPassword(c *gin.Context) {
	userName := c.Param("name")
	password := c.Param("pass")
	err := internal.ChangeUserPassword(userName, password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": fmt.Sprintf("user %s password changed", userName)})
}
