package service

import (
	"log"
	"os"
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
	DeleteProduct(id uuid.UUID) error
	EditProduct(id uuid.UUID) (model.Product, error)
	UpdateProduct(newProduct model.Product, id uuid.UUID) error
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

func (s *productService) DeleteProduct(id uuid.UUID) error {
	product, err := s.productRepository.GetProductById(id)
	if err != nil {
		log.Println("error: " + err.Error())
		return err
	}

	imageDir := "./web"
	imagePath := imageDir + product.Image

	if err := os.Remove(imagePath); err != nil {
		log.Println("error: " + err.Error())
		return err
	}

	if err := s.productRepository.DeleteProduct(product); err != nil {
		log.Println("error: " + err.Error())
		return err
	}

	return nil
}

func (s *productService) EditProduct(id uuid.UUID) (model.Product, error) {
	product, err := s.productRepository.GetProductById(id)
	if err != nil {
		log.Println("error: " + err.Error())
		return model.Product{}, err
	}

	return product, nil
}

func (s *productService) UpdateProduct(newProduct model.Product, id uuid.UUID) error {
	product, err := s.productRepository.GetProductById(id)
	if err != nil {
		log.Println("error: " + err.Error())
		return err
	}

	if newProduct.Image != "" {
		imageDir := "./web"
		imagePath := imageDir + product.Image

		if err := os.Remove(imagePath); err != nil {
			log.Println("error: " + err.Error())
			return err
		}

		product.Image = newProduct.Image
	}

	product.Name = newProduct.Name
	product.Description = newProduct.Description
	product.Quantity = newProduct.Quantity
	product.Price = newProduct.Price
	product.UpdatedAt = time.Now()

	if err := s.productRepository.UpdateProduct(product); err != nil {
		log.Println("error: " + err.Error())
		return err
	}

	return nil
}
