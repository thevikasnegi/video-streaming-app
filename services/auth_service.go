package services

import (
	"video-streaming-app/models"
	"video-streaming-app/utils"

	"gorm.io/gorm"
)

type authService struct {
	db *gorm.DB
}

func (s *authService) ValidateCredentials(user *models.User, password string) bool {
	return utils.CheckPasswordHash(password, user.Password)
}
