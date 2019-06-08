package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/amitlevy21/alpha-gopher/api"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetDevices(t *testing.T) {
	router := gin.Default()
	router.GET("/devices", api.GetDevices)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/devices", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "sda1")
	assert.Contains(t, w.Body.String(), "sda2")
	assert.Contains(t, w.Body.String(), "[SWAP]")
}

func TestBackupSingleFile(t *testing.T) {
	router := gin.Default()
	router.POST("/backup", api.Backup)
	fileName := "test.txt"
	filePath := fmt.Sprintf("/%s", fileName)
	w := httptest.NewRecorder()
	f, _ := os.Create(filePath)
	req, _ := http.NewRequest("POST", fmt.Sprintf("/backup?path=%s", filePath), nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	backedPath := fmt.Sprintf("/backup/%s", fileName)
	assert.FileExists(t, backedPath)
	f.Close()
	os.Remove(backedPath)
	os.Remove(filePath)
}

func TestBackupFolder(t *testing.T) {
	router := gin.Default()
	router.POST("/backup", api.Backup)
	folderName := "bin"
	folderPath := fmt.Sprintf("/%s", folderName)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", fmt.Sprintf("/backup?path=%s", folderPath), nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	backedPath := fmt.Sprintf("/backup/%s", folderName)
	assert.DirExists(t, backedPath)
	os.Remove(backedPath)
}
