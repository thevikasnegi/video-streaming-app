package services

import (
	"video-streaming-app/config"
)

var UserService *userService
var AuthService *authService

func InitializeServices() {
	UserService = &userService{db: config.DB}
	AuthService = &authService{db: config.DB}
}
