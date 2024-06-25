package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"kopoksu/internal/model"
)

type userRepository struct {
	DB *gorm.DB
}

type UserRepository interface {
	GetAllUser() ([]model.User, error)
	CreateUser(user model.User) error
	GetUserById(id uuid.UUID) (model.User, error)
	GetUserByEmailPassword(email, password string) (model.User, error)
	UpdateUser(user model.User) error
	DeleteUser(user model.User) error
}

func NewUserRepository(DB *gorm.DB) *userRepository {
	return &userRepository{
		DB: DB,
	}
}

func (r *userRepository) GetAllUser() ([]model.User, error) {
	var users []model.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) CreateUser(user model.User) error {
	return r.DB.Create(&user).Error
}

func (r *userRepository) GetUserById(id uuid.UUID) (model.User, error) {
	var user model.User
	if err := r.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) GetUserByEmailPassword(email, password string) (model.User, error) {
	var user model.User
	if err := r.DB.Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) UpdateUser(user model.User) error {
	return r.DB.Save(&user).Error
}

func (r *userRepository) DeleteUser(user model.User) error {
	return r.DB.Delete(&user).Error
}
