package service

import (
	"github.com/google/uuid"
	"kopoksu/internal/model"
	"kopoksu/internal/repository"
	"time"
)

type userService struct {
	userRepo repository.UserRepository
}

type UserService interface {
	GetAllUser() ([]model.UserResponse, error)
	GetUserById(id string) (model.User, error)
	CreateUser(user model.User) error
	UpdateUser(user model.User) error
	DeleteUser(id string) error
}

func NewUserService(userRepo repository.UserRepository) *userService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) GetAllUser() ([]model.UserResponse, error) {
	users, err := s.userRepo.GetAllUser()
	if err != nil {
		return nil, err
	}

	var userResponses []model.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, model.UserResponse{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return userResponses, nil
}

func (s *userService) GetUserById(id string) (model.User, error) {
	user, err := s.userRepo.GetUserById(uuid.MustParse(id))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *userService) CreateUser(user model.User) error {
	user.Id = uuid.New()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if err := s.userRepo.CreateUser(user); err != nil {
		return err
	}

	return nil
}

func (s *userService) UpdateUser(user model.User) error {
	user.UpdatedAt = time.Now()

	if err := s.userRepo.UpdateUser(user); err != nil {
		return err
	}

	return nil
}

func (s *userService) DeleteUser(id string) error {
	user, err := s.userRepo.GetUserById(uuid.MustParse(id))
	if err != nil {
		return err
	}

	if err := s.userRepo.DeleteUser(user); err != nil {
		return err
	}

	return nil
}
