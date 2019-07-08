package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/amitlevy21/alpha-gopher/api"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestExecCommand(t *testing.T) {
	url := "/terminal/exec"
	router := gin.Default()
	router.POST(url, api.ExecCommand)
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/terminal/exec?command=pwd", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "/go")
}
