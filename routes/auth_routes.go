package routes

import (
	"video-streaming-app/controllers"

	"github.com/gin-gonic/gin"
)

func InitAuthRoutes(rg *gin.RouterGroup) {
	auth := rg.Group("/auth")
	c := controllers.AuthController
	auth.POST("/login", c.Login)
}
