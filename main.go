package main

import (
	"video-streaming-app/config"
	"video-streaming-app/routes"
	"video-streaming-app/services"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	routes.RegisterUserRoutes(server)
	config.ConnectDB()
	services.InitializeServices()
	server.Run(":8080")
}
