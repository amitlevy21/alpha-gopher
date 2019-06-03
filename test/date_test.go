package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/amitlevy21/alpha-gopher/api"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetDate(t *testing.T) {
	router := gin.Default()
	router.GET("/date", api.GetDate)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/date", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), time.Now().Format("02-Jan-2006 15:04"))
}

func TestSetDate(t *testing.T) {
	// Will not work inside container. It is possible to fake it.
	router := gin.Default()
	router.POST("/date/:newDate", api.SetDate)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/date/02-Feb-2008", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), time.Now().Format("02-Jan-2006 15:04"))
}
