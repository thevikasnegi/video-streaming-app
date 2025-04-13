package controllers

import (
	"net/http"
	"video-streaming-app/responses"
	"video-streaming-app/services"
	"video-streaming-app/validations"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Login(ctx *gin.Context) {
	var req validations.LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.Error(ctx, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	if err := validate.Struct(req); err != nil {
		responses.Error(ctx, http.StatusBadRequest, "Validation error", err.Error())
		return
	}

	query := map[string]interface{}{
		"email": req.Email,
	}
	user, err := services.UserService.FindOne(query)
	if err != nil {
		responses.Error(ctx, http.StatusInternalServerError, "Error searching user", err.Error())
		return
	}

	if user.ID == uuid.Nil {
		responses.Error(ctx, http.StatusNotFound, "User not found", nil)
		return
	}

	if !services.AuthService.ValidateCredentials(user, req.Password) {
		responses.Error(ctx, http.StatusUnauthorized, "Invalid credentials", nil)
		return
	}

	responses.Success(ctx, http.StatusOK, "Logged in successfully", user)
}
