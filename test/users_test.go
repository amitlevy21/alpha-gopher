package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/amitlevy21/alpha-gopher/api"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/users", api.GetAllUsers)
	return r
}

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users", nil)
	router.ServeHTTP(w, req)

	expected := string("{\"users\":[\"root\",\"bin\",\"daemon\",\"adm\",\"lp\",\"sync\",\"shutdown\",\"halt\",\"mail\",\"operator\",\"games\",\"ftp\",\"nobody\",\"systemd-network\",\"dbus\",\"polkitd\",\"rpc\",\"unbound\",\"tss\",\"libstoragemgmt\",\"colord\",\"gluster\",\"kvmuser\",\"usbmuxd\",\"geoclue\",\"saslauth\",\"postfix\",\"rtkit\",\"pulse\",\"radvd\",\"saned\",\"chrony\",\"qemu\",\"rpcuser\",\"nfsnobody\",\"gsanslcd\",\"nm-openconnect\",\"gdm\",\"sshd\",\"avahi\",\"ntp\",\"tcpdump\",\"amit\",\"lightdm\",\"\"]}\n")

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, expected, w.Body.String())
}
