package service

import (
	"template/internal/model"
	"template/internal/repository"
)

type productService struct {
	productRepository repository.ProductRepository
}

type ProductService interface {
	GetAllProductsFormulaMilk() ([]model.Product, error)
}

func NewProductService(productRepository repository.ProductRepository) *productService {
	return &productService{
		productRepository: productRepository,
	}
}

func (s *productService) GetAllProductsFormulaMilk() ([]model.Product, error) {
	return s.productRepository.GetAllProductsFormulaMilk()
}
