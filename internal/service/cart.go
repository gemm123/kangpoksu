package service

import (
	"kopoksu/internal/model"
	"kopoksu/internal/repository"
	"log"
)

type cartService struct {
	productRepository repository.ProductRepository
}

type CartService interface {
	GetAccumulationTotalCart(cart []model.Cart) (int, error)
}

func NewCartService(productRepository repository.ProductRepository) *cartService {
	return &cartService{
		productRepository: productRepository,
	}
}

func (s *cartService) GetAccumulationTotalCart(cart []model.Cart) (int, error) {
	var totalOrder int

	for i := 0; i < len(cart); i++ {
		product, err := s.productRepository.GetProductById(cart[i].Id)
		if err != nil {
			log.Println("error: " + err.Error())
			return 0, err
		}
		cart[i].Name = product.Name
		cart[i].Image = product.Image
		cart[i].Total = product.Price * cart[i].Amount
		totalOrder = totalOrder + cart[i].Total
	}

	return totalOrder, nil
}
