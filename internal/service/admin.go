package service

import (
	"kopoksu/config"
	"kopoksu/internal/repository"
	"log"
)

type adminService struct {
	userRepo repository.UserRepository
}

type AdminService interface {
	CheckCredentials(email, password string) bool
}

func NewAdminService(userRepo repository.UserRepository) *adminService {
	return &adminService{
		userRepo: userRepo,
	}
}

func (s *adminService) CheckCredentials(email, password string) bool {
	masterEmail := config.MasterEmail()
	masterPassword := config.MasterPassword()

	if email == masterEmail && password == masterPassword {
		return true
	}

	user, err := s.userRepo.GetUserByEmailPassword(email, password)
	if err != nil {
		log.Println("error when get user by email and password: ", err.Error())
		return false
	}

	if user.Name != "" {
		return true
	}

	log.Println("wrong credentials")

	return false
}
