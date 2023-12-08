package repository

import (
	"template/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type productRepository struct {
	DB *gorm.DB
}

type ProductRepository interface {
	GetAllProductsFormulaMilk() ([]model.Product, error)
	SaveProductFormulaMilk(formulaMilk model.Product) error
	DeleteProduct(product model.Product) error
	GetProductById(id uuid.UUID) (model.Product, error)
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

func (r *productRepository) SaveProductFormulaMilk(formulaMilk model.Product) error {
	err := r.DB.Create(&formulaMilk).Error
	return err
}

func (r *productRepository) DeleteProduct(product model.Product) error {
	return r.DB.Delete(&product).Error
}

func (r *productRepository) GetProductById(id uuid.UUID) (model.Product, error) {
	var product model.Product
	err := r.DB.First(&product, "id = ?", id).Error
	return product, err
}
