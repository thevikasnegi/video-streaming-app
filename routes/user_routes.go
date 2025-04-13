package routes

import (
	"video-streaming-app/controllers"

	"github.com/gin-gonic/gin"
)

func InitUserRoutes(rg *gin.RouterGroup) {
	c := controllers.UserController
	user := rg.Group("/users")
	user.POST("/", c.CreateUser)
	user.GET("/", c.GetAllUsers)
	user.GET("/:id", c.GetUser)
	user.PUT("/:id", c.UpdateUser)
	user.DELETE("/:id", c.DeleteUser)
}
