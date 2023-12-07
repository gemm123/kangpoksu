package service

import (
	"template/internal/model"
	"template/internal/repository"
	"time"

	"github.com/google/uuid"
)

type productService struct {
	productRepository repository.ProductRepository
}

type ProductService interface {
	GetAllProductsFormulaMilk() ([]model.Product, error)
	SaveProductFormulaMilk(formulaMilk model.Product) error
}

func NewProductService(productRepository repository.ProductRepository) *productService {
	return &productService{
		productRepository: productRepository,
	}
}

func (s *productService) GetAllProductsFormulaMilk() ([]model.Product, error) {
	return s.productRepository.GetAllProductsFormulaMilk()
}

func (s *productService) SaveProductFormulaMilk(formulaMilk model.Product) error {
	formulaMilk.Id = uuid.New()
	formulaMilk.CategoryId = uuid.MustParse("ea600c63-283a-415e-8ed1-b10d12c544a0")
	formulaMilk.CreatedAt = time.Now()
	formulaMilk.UpdatedAt = time.Now()

	return s.productRepository.SaveProductFormulaMilk(formulaMilk)
}
