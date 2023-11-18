package repository

import "gorm.io/gorm"

type repository struct {
	DB *gorm.DB
}

type Repository interface {
	GetMessage() string
}

func NewRepository(DB *gorm.DB) *repository {
	return &repository{DB: DB}
}

func (r *repository) GetMessage() string {
	return "Hello World"
}
