package main

import (
	"github.com/amitlevy21/alpha-gopher/api"
	"github.com/gin-contrib/cors"
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
	router.POST("/filesystem/rm", api.Remove)
	router.GET("/filesystem/ls/all", api.ListAll)
	router.GET("/owner", api.GetOwner)
	router.POST("/owner/:newOwner", api.ChangeOwnership)
	router.GET("/permission", api.GetCurrentPermissions)
	router.GET("/users/all", api.GetAllUsers)
	router.POST("/users/new/:name", api.AddUser)
	router.DELETE("/users/:name", api.DeleteUser)
	router.POST("/users/cred/:name/:pass", api.ChangeUserPassword)
	router.POST("/terminal/exec", api.ExecCommand)
	router.GET("/find", api.Find)

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://172.20.128.1:8080"}
	router.Use(cors.New(config))
	router.Run(":3000")
}
