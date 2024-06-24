package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id        uuid.UUID `gorm:"id"`
	Name      string    `gorm:"name"`
	Email     string    `gorm:"email"`
	Password  string    `gorm:"password"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}

type UserResponse struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}
