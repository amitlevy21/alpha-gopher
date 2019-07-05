package main

import (
	"github.com/amitlevy21/alpha-gopher/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/devices", api.GetDevices)
	router.POST("/backup", api.Backup)
	router.GET("/date", api.GetDate)
	router.POST("/date/:newDate", api.SetDate)
	router.POST("/filesystem/cp", api.Copy)
	router.POST("/filesystem/ls", api.List)
	router.POST("/filesystem/mv", api.Move)
	router.GET("/owner", api.GetOwner)
	router.POST("/owner/:newOwner", api.ChangeOwnership)
	router.GET("/permission", api.GetCurrentPermissions)

	router.Run(":3000")
}
