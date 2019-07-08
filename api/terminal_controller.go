package api

import (
	"net/http"
	"strings"

	"github.com/amitlevy21/alpha-gopher/internal"

	"github.com/gin-gonic/gin"
)

func ExecCommand(c *gin.Context) {
	cmd := c.Query("command")
	cmdSplt := strings.Split(cmd, " ")
	std, err := internal.ExecCommand(cmdSplt)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "std": std})
		return
	}
	c.JSON(http.StatusOK, gin.H{"std": std})
}
