package controllers

import (
	"net/http"
	"video-streaming-app/models"
	"video-streaming-app/utils"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	var payload *models.UserRequest
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse the request body."})
		return
	}
	user := models.User{
		FirstName:    payload.FirstName,
		LastName:     payload.LastName,
		Email:        payload.Email,
		MobileNumber: payload.MobileNumber,
		Password:     hashedPassword,
	}
	user.Save()
	ctx.JSON(http.StatusOK, gin.H{"message": "User created successfully", "data": user})
}
