package controllers

import (
	"errors"
	"net/http"
	"video-streaming-app/models"
	"video-streaming-app/responses"
	"video-streaming-app/services"
	"video-streaming-app/validations"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var validate = validator.New()

func CreateUser(ctx *gin.Context) {
	var req validations.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.Error(ctx, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	if err := validate.Struct(req); err != nil {
		responses.Error(ctx, http.StatusBadRequest, "Validation error", err.Error())
		return
	}
	queryMap := map[string]interface{}{
		"email":         req.Email,
		"mobile_number": req.MobileNumber,
	}
	existingUser, err := services.UserService.FindOne(queryMap)
	if err != nil {
		responses.Error(ctx, http.StatusInternalServerError, "Error searching user", err.Error())
		return
	}
	if existingUser.ID != uuid.Nil {
		responses.Error(ctx, http.StatusConflict,
			"User Already exists",
			errors.New("a user already exists with same email or mobile number"))
		return
	}
	user := &models.User{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Email:        req.Email,
		MobileNumber: req.MobileNumber,
		Password:     req.Password,
	}

	if err := services.UserService.Create(user); err != nil {
		responses.Error(ctx, http.StatusInternalServerError, "Failed to create user", err.Error())
		return
	}

	responses.Success(ctx, http.StatusCreated, "User created successfully", user)
}

func GetUser(ctx *gin.Context) {
	idParam := ctx.Param("id")
	userID, err := uuid.Parse(idParam)
	if err != nil {
		responses.Error(ctx, http.StatusBadRequest, "Invalid UUID format", nil)
		return
	}

	user, err := services.UserService.FindByID(userID)
	if err != nil {
		responses.Error(ctx, http.StatusNotFound, "User not found", err.Error())
		return
	}

	responses.Success(ctx, http.StatusOK, "User retrieved successfully", user)
}

func GetAllUsers(ctx *gin.Context) {
	users, err := services.UserService.FindAll()
	if err != nil {
		responses.Error(ctx, http.StatusInternalServerError, "Error fetching users", nil)
		return
	}
	responses.Success(ctx, http.StatusOK, "User retrieved successfully", users)
}

func UpdateUser(ctx *gin.Context) {
	idParam := ctx.Param("id")
	userID, err := uuid.Parse(idParam)
	if err != nil {
		responses.Error(ctx, http.StatusBadRequest, "Invalid UUID format", nil)
		return
	}

	var updateData map[string]interface{}
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		responses.Error(ctx, http.StatusBadRequest, "Invalid update data", err.Error())
		return
	}

	if err := services.UserService.Update(userID, updateData); err != nil {
		responses.Error(ctx, http.StatusInternalServerError, "Failed to update user", err.Error())
		return
	}

	responses.Success(ctx, http.StatusOK, "User updated", nil)
}

func DeleteUser(ctx *gin.Context) {
	idParam := ctx.Param("id")
	userID, err := uuid.Parse(idParam)
	if err != nil {
		responses.Error(ctx, http.StatusBadRequest, "Invalid UUID format", nil)
		return
	}

	if err := services.UserService.Delete(userID); err != nil {
		responses.Error(ctx, http.StatusInternalServerError, "Failed to delete user", err.Error())
		return
	}

	responses.Success(ctx, http.StatusOK, "User deleted", nil)
}
