package repository

import (
	"template/internal/model"

	"gorm.io/gorm"
)

type productRepository struct {
	DB *gorm.DB
}

type ProductRepository interface {
	GetAllProductsFormulaMilk() ([]model.Product, error)
}

func NewProductRepository(DB *gorm.DB) *productRepository {
	return &productRepository{
		DB: DB,
	}
}

func (r *productRepository) GetAllProductsFormulaMilk() ([]model.Product, error) {
	var products []model.Product
	err := r.DB.Where("category_id = 'ea600c63-283a-415e-8ed1-b10d12c544a0'").Find(&products).Error
	return products, err
}
