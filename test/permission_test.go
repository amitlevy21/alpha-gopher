package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/amitlevy21/alpha-gopher/api"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetOwner(t *testing.T) {
	router := gin.Default()
	router.GET("/owner", api.GetOwner)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/owner?path=/etc/passwd", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "root")
}

func TestChangeOwnership(t *testing.T) {
	router := gin.Default()
	router.POST("/owner/:newOwner", api.ChangeOwnership)
	router.GET("/owner", api.GetOwner)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/owner/nobody?path=/etc/passwd", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	req, _ = http.NewRequest("GET", "/owner?path=/etc/passwd", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "nobody")

	req, _ = http.NewRequest("POST", "/owner/root?path=/etc/passwd", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestGetCurrnetPermissions(t *testing.T) {
	router := gin.Default()
	router.GET("/permission", api.GetCurrentPermissions)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/permission?path=/etc/passwd", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "644")
}

// func TestChangePermissions(t *testing.T) {
// 	router := gin.Default()
// 	router.POST("/permission/:owner", api.ChangeOwnership)
// 	router.GET("/permission", api.GetOwner)

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("POST", "/permission/nobody?path=/etc/passwd", nil)
// 	router.ServeHTTP(w, req)
// 	assert.Equal(t, 200, w.Code)

// 	req, _ = http.NewRequest("GET", "/permission?path=/etc/passwd", nil)
// 	router.ServeHTTP(w, req)
// 	assert.Equal(t, 200, w.Code)
// 	assert.Contains(t, w.Body.String(), "nobody")

// 	req, _ = http.NewRequest("POST", "/permission/root?path=/etc/passwd", nil)
// 	router.ServeHTTP(w, req)
// 	assert.Equal(t, 200, w.Code)
// }
