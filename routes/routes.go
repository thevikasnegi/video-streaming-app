package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine) {
	rg := router.Group("/api")
	InitAuthRoutes(rg)
	InitUserRoutes(rg)
}
