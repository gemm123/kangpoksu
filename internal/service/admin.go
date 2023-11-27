package service

import (
	"log"
	"template/config"
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

	if email != adminEmail || password != adminPassword {
		log.Println("wrong credentials")
		return false
	}

	return true
}
