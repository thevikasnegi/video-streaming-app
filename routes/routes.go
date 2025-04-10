package routes

import (
	"video-streaming-app/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/video", getVideos)
	server.POST("/api/user", controllers.CreateUser)
}
