package services

import (
	"errors"
	"fmt"
	"strings"
	"video-streaming-app/models"
	"video-streaming-app/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var ErrDuplicateUser = errors.New("user already exists with the same email or mobile number")

type userService struct {
	db *gorm.DB
}

func (s *userService) Create(user *models.User) error {
	user.ID = uuid.New()
	hashedPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashedPassword
	return s.db.Create(&user).Error
}

func (s *userService) FindByID(id uuid.UUID) (*models.User, error) {
	var user models.User
	err := s.db.First(&user, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

func (s *userService) FindAll() ([]models.User, error) {
	var users []models.User
	err := s.db.Find(&users).Error
	return users, err
}

func (s *userService) Update(id uuid.UUID, updatedData map[string]interface{}) error {
	params := make(map[string]interface{})
	if val, ok := updatedData["email"]; ok && val != nil {
		params["email"] = val
	}
	if val, ok := updatedData["mobile_number"]; ok && val != nil {
		params["mobile_number"] = val
	}

	if len(params) != 0 {
		existingUser, err := s.FindOne(params)
		if err != nil {
			return err
		}
		if existingUser != nil && existingUser.ID != uuid.Nil && existingUser.ID != id {
			return ErrDuplicateUser
		}
	}

	return s.db.Model(&models.User{}).Where("id = ?", id).Updates(updatedData).Error
}

func (s *userService) Delete(id uuid.UUID) error {
	return s.db.Delete(&models.User{}, id).Error
}

func (s *userService) Search(params map[string]interface{}) ([]models.User, error) {
	var users []models.User
	var conditions []string
	var values []interface{}

	for k, v := range params {
		conditions = append(conditions, fmt.Sprintf("%s = ?", k))
		values = append(values, v)
	}

	whereClause := strings.Join(conditions, " OR ")

	query := s.db.Where(whereClause, values...)
	err := query.Find(&users).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return users, nil
	}
	return users, err
}

func (s *userService) FindOne(params map[string]interface{}) (*models.User, error) {
	var user models.User
	var conditions []string
	var values []interface{}

	for k, v := range params {
		conditions = append(conditions, fmt.Sprintf("%s = ?", k))
		values = append(values, v)
	}

	whereClause := strings.Join(conditions, " OR ")
	err := s.db.Where(whereClause, values...).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &user, nil
	}
	return &user, err
}
