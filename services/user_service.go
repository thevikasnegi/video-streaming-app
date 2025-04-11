package services

import (
	"errors"
	"fmt"
	"strings"
	"video-streaming-app/config"
	"video-streaming-app/models"
	"video-streaming-app/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateUser(user *models.User) error {
	user.ID = uuid.New()
	hashedPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashedPassword
	return config.DB.Create(&user).Error
}

func GetUserByID(id uuid.UUID) (*models.User, error) {
	var user models.User
	err := config.DB.First(&user, "id = ?", id).Error
	return &user, err
}

func GetUsers() ([]models.User, error) {
	var users []models.User
	err := config.DB.Find(&users).Error
	return users, err
}

func UpdateUser(id uuid.UUID, updatedData map[string]interface{}) error {
	return config.DB.Model(&models.User{}).Where("id = ?", id).Updates(updatedData).Error
}

func DeleteUser(id uuid.UUID) error {
	return config.DB.Delete(&models.User{}, id).Error
}

func SearchUser(params map[string]interface{}, findOne bool) (users []models.User, err error) {
	var conditions []string
	var values []interface{}
	for k, v := range params {
		conditions = append(conditions, fmt.Sprintf("%s = ?", k))
		values = append(values, v)
	}
	whereClause := strings.Join(conditions, " OR ")
	err = config.DB.Where(whereClause, values...).Find(&users).Error
	return users, err
}

func FindUser(params map[string]interface{}) (*models.User, error) {
	var user models.User
	var conditions []string
	var values []interface{}
	for k, v := range params {
		conditions = append(conditions, fmt.Sprintf("%s = ?", k))
		values = append(values, v)
	}
	whereClause := strings.Join(conditions, " OR ")
	err := config.DB.Where(whereClause, values...).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &user, nil
	}
	return &user, err
}
