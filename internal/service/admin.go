package service

import (
	"kopoksu/config"
	"log"
)

type adminService struct {
}

type AdminService interface {
	CheckCredentials(email, password string) bool
}

func NewAdminService() *adminService {
	return &adminService{}
}

func (s *adminService) CheckCredentials(email, password string) bool {
	adminEmail := config.AdminEmail()
	adminPassword := config.AdminPassword()
	masterEmail := config.MasterEmail()
	masterPassword := config.MasterPassword()

	if email == adminEmail && password == adminPassword {
		return true
	} else if email == masterEmail && password == masterPassword {
		return true
	}

	log.Println("wrong credentials")

	return false
}
