package main

import (
	"github.com/amitlevy21/alpha-gopher/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/devices", api.GetDevices)

	router.Run()
}
