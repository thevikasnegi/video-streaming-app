package responses

import (
	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(ctx *gin.Context, code int, message string, data interface{}) {
	ctx.JSON(code, APIResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func Error(ctx *gin.Context, code int, message string, data interface{}) {
	ctx.JSON(code, APIResponse{
		Status:  "error",
		Message: message,
		Data:    data,
	})
}
