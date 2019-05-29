package api

import (
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {

	users, err := exec.Command("awk -F: '{ print $1}' /etc/passwd").Output()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err})
	}
	c.JSON(200, gin.H{"message": users})
}
