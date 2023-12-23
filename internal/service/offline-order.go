package service

import (
	"github.com/google/uuid"
	"kopoksu/internal/model"
	"kopoksu/internal/repository"
	"log"
	"time"
)

type offlineOrderService struct {
	orderRepository   repository.OfflineOrderRepository
	productRepository repository.ProductRepository
}

type OfflineOrderService interface {
	SaveOfflineOrder(offlineOrder model.OfflineOrder, cart []model.Cart) error
	GetAllOfflineOrder() ([]model.OfflineOrder, error)
	EditOfflineOrder(id uuid.UUID) (model.EditOfflineOrderResponse, error)
	UpdateStatusOfflineOrder(id uuid.UUID, status string) error
	DeleteOfflineOrder(id uuid.UUID) error
}

func NewOfflineOrderService(offlineOrderRepository repository.OfflineOrderRepository, productRepository repository.ProductRepository) *offlineOrderService {
	return &offlineOrderService{
		orderRepository:   offlineOrderRepository,
		productRepository: productRepository,
	}
}

func (s *offlineOrderService) SaveOfflineOrder(offlineOrder model.OfflineOrder, cart []model.Cart) error {
	offlineOrder.Id = uuid.New()
	offlineOrder.CreatedAt = time.Now()
	offlineOrder.UpdatedAt = time.Now()
	offlineOrder.Status = "Menunggu konfirmasi pembayaran"

	if err := s.orderRepository.SaveOfflineOrder(offlineOrder); err != nil {
		log.Println("error: " + err.Error())
		return err
	}

	for _, c := range cart {
		detailOrder := model.DetailOfflineOrder{
			Id:             uuid.New(),
			OfflineOrderId: offlineOrder.Id,
			ProductId:      c.Id,
			Amount:         c.Amount,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}

		if err := s.orderRepository.SaveDetailOfflineOrder(detailOrder); err != nil {
			log.Println("error: " + err.Error())
			return err
		}

		product, err := s.productRepository.GetProductById(c.Id)
		if err != nil {
			log.Println("error: " + err.Error())
			return err
		}

		product.Quantity = product.Quantity - c.Amount

		if err := s.productRepository.UpdateProduct(product); err != nil {
			log.Println("error: " + err.Error())
			return err
		}
	}

	return nil
}

func (s *offlineOrderService) GetAllOfflineOrder() ([]model.OfflineOrder, error) {
	offlineOrders, err := s.orderRepository.GetAllOfflineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return offlineOrders, err
	}

	return offlineOrders, nil
}

func (s *offlineOrderService) EditOfflineOrder(id uuid.UUID) (model.EditOfflineOrderResponse, error) {
	var editOfflineOrderResponse model.EditOfflineOrderResponse

	offlineOrder, err := s.orderRepository.GetOfflineOrderById(id)
	if err != nil {
		log.Println("error: " + err.Error())
		return editOfflineOrderResponse, err
	}

	detailOfflineOrderResponse, err := s.orderRepository.GetDetailOfflineOrderByOfflineOrderId(id)
	if err != nil {
		log.Println("error: " + err.Error())
		return editOfflineOrderResponse, err
	}

	editOfflineOrderResponse.Id = id
	editOfflineOrderResponse.Name = offlineOrder.Name
	editOfflineOrderResponse.PhoneNumber = offlineOrder.PhoneNumber
	editOfflineOrderResponse.Total = offlineOrder.Total
	editOfflineOrderResponse.Status = offlineOrder.Status
	editOfflineOrderResponse.DetailOfflineOrderResponse = detailOfflineOrderResponse

	return editOfflineOrderResponse, nil
}

func (s *offlineOrderService) UpdateStatusOfflineOrder(id uuid.UUID, status string) error {
	data := map[string]interface{}{
		"status":     status,
		"updated_at": time.Now(),
	}
	err := s.orderRepository.UpdateOfflineOrder(id, data)
	if err != nil {
		log.Println("error: " + err.Error())
		return err
	}

	return nil
}

func (s *offlineOrderService) DeleteOfflineOrder(id uuid.UUID) error {
	offlineOrder, err := s.orderRepository.GetOfflineOrderById(id)
	if err != nil {
		log.Println("error: " + err.Error())
		return err
	}

	detailOfflineOrders, err := s.orderRepository.GetAllDetailOfflineOrderByOfflineOrderId(id)
	if err != nil {
		log.Println("error: " + err.Error())
		return err
	}

	for _, doo := range detailOfflineOrders {
		if err := s.orderRepository.DeleteDetailOfflineOrder(doo); err != nil {
			log.Println("error: " + err.Error())
			return err
		}
	}

	if err := s.orderRepository.DeleteOfflineOrder(offlineOrder); err != nil {
		log.Println("error: " + err.Error())
		return err
	}

	return nil
}
