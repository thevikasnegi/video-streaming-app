package main

import (
	"video-streaming-app/config"
	"video-streaming-app/models"
	"video-streaming-app/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	server := gin.Default()
	_ = godotenv.Load("config/.env")
	routes.RegisterRoutes(server)
	config.ConnectDB()
	err := config.DB.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		panic("Failed to auto-migrate database: " + err.Error())
	}
	server.Run(":8080")
}
