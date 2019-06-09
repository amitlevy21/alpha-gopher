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

func TestCopyFile(t *testing.T) {
	router := gin.Default()
	router.POST("/fs/cp", api.Copy)
	srcFile := "test.txt"
	srcPath := fmt.Sprintf("/%s", srcFile)
	dstFile := "test.txt"
	dstPath := fmt.Sprintf("/bin/%s", dstFile)
	w := httptest.NewRecorder()

	os.Create(srcPath)
	req, _ := http.NewRequest("POST", fmt.Sprintf("/fs/cp?src=%s&dst=%s&opts=-r", srcPath, dstPath), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.FileExists(t, dstPath)

	os.Remove(srcPath)
	os.Remove(dstPath)
}

func TestCopyWithoutOpts(t *testing.T) {
	router := gin.Default()
	router.POST("/fs/cp", api.Copy)
	srcFile := "test.txt"
	srcPath := fmt.Sprintf("/%s", srcFile)
	dstFile := "test.txt"
	dstPath := fmt.Sprintf("/bin/%s", dstFile)
	w := httptest.NewRecorder()

	os.Create(srcPath)
	req, _ := http.NewRequest("POST", fmt.Sprintf("/fs/cp?src=%s&dst=%s", srcPath, dstPath), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.FileExists(t, dstPath)

	os.Remove(srcPath)
	os.Remove(dstPath)
}

func TestMoveFile(t *testing.T) {
	router := gin.Default()
	router.POST("/fs/mv", api.Move)
	srcFile := "test.txt"
	srcPath := fmt.Sprintf("/%s", srcFile)
	dstFile := "test.txt"
	dstPath := fmt.Sprintf("/bin/%s", dstFile)
	w := httptest.NewRecorder()

	os.Create(srcPath)
	req, _ := http.NewRequest("POST", fmt.Sprintf("/fs/mv?src=%s&dst=%s", srcPath, dstPath), nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.FileExists(t, dstPath)

	os.Remove(srcPath)
	os.Remove(dstPath)
}

func TestRenameFile(t *testing.T) {
	router := gin.Default()
	router.POST("/fs/mv", api.Move)
	srcFile := "test.txt"
	srcPath := fmt.Sprintf("/%s", srcFile)
	dstFile := "rename.txt"
	dstPath := fmt.Sprintf("/%s", dstFile)
	w := httptest.NewRecorder()

	os.Create(srcPath)
	req, _ := http.NewRequest("POST", fmt.Sprintf("/fs/mv?src=%s&dst=%s", srcPath, dstPath), nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.FileExists(t, dstPath)

	os.Remove(srcPath)
	os.Remove(dstPath)
}

func TestListFiles(t *testing.T) {
	router := gin.Default()
	router.POST("/fs/ls", api.List)
	path := "/lib"
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", fmt.Sprintf("/fs/ls?path=%s", path), nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "systemd")
}

func TestListFilesOpts(t *testing.T) {
	router := gin.Default()
	router.POST("/fs/ls", api.List)
	path := "/lib"
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", fmt.Sprintf("/fs/ls?path=%s&opts=-lst", path), nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "systemd")
	assert.Contains(t, w.Body.String(), "root")
}
