package service

import (
	"kopoksu/helper"
	"kopoksu/internal/model"
	"kopoksu/internal/repository"
	"log"
	"os"
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
	GetAllProductsBabyDiaper() ([]model.Product, error)
	SaveProductBabyDiaper(babyDiaper model.Product) error
	GetAllProductsAdultDiaper() ([]model.Product, error)
	SaveProductAdultDiaper(adultDiaper model.Product) error
	GetAllProductsFormulaMilkLimit(limit int) ([]model.Product, error)
	GetAllProductsBabyDiaperLimit(limit int) ([]model.Product, error)
	GetAllProductsAdultDiaperLimit(limit int) ([]model.Product, error)
	GetProductById(id uuid.UUID) (model.Product, error)
	SearchProductsByName(name string) ([]model.SearchResult, error)
}

func NewProductService(productRepository repository.ProductRepository) *productService {
	return &productService{
		productRepository: productRepository,
	}
}

func (s *productService) GetAllProductsFormulaMilk() ([]model.Product, error) {
	products, err := s.productRepository.GetAllProductsFormulaMilk()
	if err != nil {
		log.Println("error: " + err.Error())
		return products, err
	}

	return products, nil
}

func (s *productService) GetAllProductsFormulaMilkLimit(limit int) ([]model.Product, error) {
	products, err := s.productRepository.GetAllProductsFormulaMilkLimit(limit)
	if err != nil {
		log.Println("error: " + err.Error())
		return products, err
	}

	for i := 0; i < len(products); i++ {
		products[i].PriceFormatted = helper.FormatRupiah(float64(products[i].Price))
	}

	return products, nil
}

func (s *productService) SaveProductFormulaMilk(formulaMilk model.Product) error {
	formulaMilk.Id = uuid.New()
	formulaMilk.CategoryId = uuid.MustParse("ea600c63-283a-415e-8ed1-b10d12c544a0")
	formulaMilk.CreatedAt = time.Now()
	formulaMilk.UpdatedAt = time.Now()

	return s.productRepository.SaveProduct(formulaMilk)
}

func (s *productService) GetAllProductsBabyDiaper() ([]model.Product, error) {
	diapers, err := s.productRepository.GetAllProductsBabyDiaper()
	if err != nil {
		log.Println("error: " + err.Error())
		return diapers, err
	}

	return diapers, nil
}

func (s *productService) GetAllProductsBabyDiaperLimit(limit int) ([]model.Product, error) {
	diapers, err := s.productRepository.GetAllProductsBabyDiaperLimit(limit)
	if err != nil {
		log.Println("error: " + err.Error())
		return diapers, err
	}

	for i := 0; i < len(diapers); i++ {
		diapers[i].PriceFormatted = helper.FormatRupiah(float64(diapers[i].Price))
	}

	return diapers, nil
}

func (s *productService) SaveProductBabyDiaper(babyDiaper model.Product) error {
	babyDiaper.Id = uuid.New()
	babyDiaper.CategoryId = uuid.MustParse("981464fb-3241-4a33-97ae-33b110e2d4aa")
	babyDiaper.CreatedAt = time.Now()
	babyDiaper.UpdatedAt = time.Now()

	if err := s.productRepository.SaveProduct(babyDiaper); err != nil {
		log.Println("error: " + err.Error())
		return err
	}

	return nil
}

func (s *productService) GetAllProductsAdultDiaper() ([]model.Product, error) {
	diapers, err := s.productRepository.GetAllProductsAdultDiaper()
	if err != nil {
		log.Println("error: " + err.Error())
		return diapers, err
	}

	return diapers, nil
}

func (s *productService) GetAllProductsAdultDiaperLimit(limit int) ([]model.Product, error) {
	diapers, err := s.productRepository.GetAllProductsAdultDiaperLimit(limit)
	if err != nil {
		log.Println("error: " + err.Error())
		return diapers, err
	}

	for i := 0; i < len(diapers); i++ {
		diapers[i].PriceFormatted = helper.FormatRupiah(float64(diapers[i].Price))
	}

	return diapers, nil
}

func (s *productService) GetProductById(id uuid.UUID) (model.Product, error) {
	product, err := s.productRepository.GetProductById(id)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *productService) SaveProductAdultDiaper(adultDiaper model.Product) error {
	adultDiaper.Id = uuid.New()
	adultDiaper.CategoryId = uuid.MustParse("f5976ce9-7496-4fd2-8322-3beaef36e4d8")
	adultDiaper.CreatedAt = time.Now()
	adultDiaper.UpdatedAt = time.Now()

	if err := s.productRepository.SaveProduct(adultDiaper); err != nil {
		log.Println("error: " + err.Error())
		return err
	}

	return nil
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

func (s *productService) SearchProductsByName(name string) ([]model.SearchResult, error) {
	searchResult, err := s.productRepository.SearchProductsByName(name)
	if err != nil {
		log.Println("error: " + err.Error())
		return searchResult, err
	}

	return searchResult, nil
}
