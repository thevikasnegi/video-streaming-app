package routes

import (
	"video-streaming-app/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine) {
	userGroup := router.Group("/api/users")
	{
		userGroup.POST("/", controllers.UserController.CreateUser)
		userGroup.GET("/", controllers.UserController.GetAllUsers)
		userGroup.GET("/:id", controllers.UserController.GetUser)
		userGroup.PUT("/:id", controllers.UserController.UpdateUser)
		userGroup.DELETE("/:id", controllers.UserController.DeleteUser)
	}

	authGroup := router.Group("/api/auth")
	{
		authGroup.POST("/login", controllers.Login)
	}

}
