package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getVideos(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "route working"})
}
