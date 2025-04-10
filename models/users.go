package models

import (
	"time"
	"video-streaming-app/config"
	"video-streaming-app/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	FirstName    string         `json:"first_name"`
	LastName     string         `json:"last_name"`
	Email        string         `gorm:"uniqueIndex;not null" json:"email"`
	MobileNumber string         `gorm:"uniqueIndex" json:"mobile_number"`
	Password     string         `gorm:"not null" json:"-"`
	IsDeleted    bool           `gorm:"default:false" json:"-"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

type UserRequest struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	MobileNumber string `json:"mobile_number"`
	Password     string `json:"password"`
}

func (u *User) SoftDelete() *User {
	u.IsDeleted = true
	u.DeletedAt.Time = time.Now()
	u.DeletedAt.Valid = true
	return u
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	hashedPassword, _ := utils.HashPassword(u.Password)
	u.Password = hashedPassword
	return nil
}

func (u *User) Save() (err error) {
	config.DB.Create(&u)
	return
}
