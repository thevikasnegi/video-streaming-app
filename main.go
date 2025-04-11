package main

import (
	"video-streaming-app/config"
	"video-streaming-app/models"
	"video-streaming-app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	routes.RegisterUserRoutes(server)
	config.ConnectDB()
	err := config.DB.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		panic("Failed to auto-migrate database: " + err.Error())
	}
	server.Run(":8080")
}
