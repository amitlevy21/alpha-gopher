package api

import (
	"net/http"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	users, err := exec.Command("cut", "-d:", "-f1", "/etc/passwd").Output()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err})
	}
	usersArr := strings.Split(string(users), "\n")

	c.JSON(http.StatusOK, gin.H{"users": usersArr})
}
