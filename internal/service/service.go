package service

import "template/internal/repository"

type service struct {
	repository repository.Repository
}

type Service interface {
	GetMessage() string
}

func NewService(repository repository.Repository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetMessage() string {
	message := s.repository.GetMessage()
	return message
}
