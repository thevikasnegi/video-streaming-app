package services

import (
	"video-streaming-app/config"
)

var UserService *userService

func InitializeServices() {
	UserService = &userService{db: config.DB}
}
