package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/amitlevy21/alpha-gopher/api"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetAllUsers(t *testing.T) {
	url := "/users"
	router := gin.Default()
	router.GET(url, api.GetAllUsers)
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", url, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "root")
}

func TestAddNonExistingUser(t *testing.T) {
	router := gin.Default()
	router.POST("/users/new/:name", api.AddUser)
	router.DELETE("/users/:name", api.DeleteUser)
	router.GET("/users", api.GetAllUsers)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users/new/i_dont_exist", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "user i_dont_exist created")
	req, _ = http.NewRequest("GET", "/users", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "i_dont_exist")

	req, _ = http.NewRequest("DELETE", "/users/i_dont_exist", nil)
	router.ServeHTTP(w, req)
}

func TestDeleteExistingUser(t *testing.T) {
	router := gin.Default()
	router.POST("/users/new/:name", api.AddUser)
	router.DELETE("/users/:name", api.DeleteUser)
	router.GET("/users", api.GetAllUsers)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users/new/bla_bla", nil)
	router.ServeHTTP(w, req)
	req, _ = http.NewRequest("GET", "/users", nil)
	router.ServeHTTP(w, req)
	req, _ = http.NewRequest("DELETE", "/users/bla_bla", nil)
	router.ServeHTTP(w, req)
	req, _ = http.NewRequest("GET", "/users", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.NotContains(t, "bla_bla", w.Body.String())
}

func TestChangeUserPassword(t *testing.T) {
	router := gin.Default()
	router.POST("/users/new/:name", api.AddUser)
	router.POST("/users/cred/:name/:pass", api.ChangeUserPassword)
	router.DELETE("/users/:name", api.DeleteUser)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users/new/jhon", nil)
	router.ServeHTTP(w, req)
	req, _ = http.NewRequest("POST", "/users/cred/jhon/mypass", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "user jhon password changed")

	req, _ = http.NewRequest("DELETE", "/users/jhon", nil)
	router.ServeHTTP(w, req)
}
