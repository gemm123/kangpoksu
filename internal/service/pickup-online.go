package service

import (
	"github.com/google/uuid"
	"kopoksu/helper"
	"kopoksu/internal/model"
	"kopoksu/internal/repository"
	"log"
	"time"
)

type pickupOnlineOrderService struct {
	pickupOnlineOrderRepository repository.PickupOnlineOrderRepository
	productRepository           repository.ProductRepository
}

type PickupOnlineOrderService interface {
	SavePickupOnlineOrder(pickupOnlineOrder model.PickupOnlineOrder, cart []model.Cart) error
	GetAllPickupOnlineOrder() ([]model.PickupOnlineOrder, error)
	EditPickupOnlineOrder(id uuid.UUID) (model.EditPickupOnlineOrderResponse, error)
	UpdateStatusPickupOnlineOrder(id uuid.UUID, status string) error
	DeletePickupOnlineOrder(id uuid.UUID) error
	CountPickupOnlineOrderByStatus(status string) (int, error)
}

func NewPickupOnlineOrderService(pickupOnlineOrderRepository repository.PickupOnlineOrderRepository, productRepository repository.ProductRepository) *pickupOnlineOrderService {
	return &pickupOnlineOrderService{
		pickupOnlineOrderRepository: pickupOnlineOrderRepository,
		productRepository:           productRepository,
	}
}

func (s *pickupOnlineOrderService) SavePickupOnlineOrder(pickupOnlineOrder model.PickupOnlineOrder, cart []model.Cart) error {
	pickupOnlineOrder.Id = uuid.New()
	pickupOnlineOrder.CreatedAt = time.Now()
	pickupOnlineOrder.UpdatedAt = time.Now()
	pickupOnlineOrder.Status = "Menunggu konfirmasi pembayaran"

	if err := s.pickupOnlineOrderRepository.SavePickupOnlineOrder(pickupOnlineOrder); err != nil {
		log.Println("error: " + err.Error())
		return err
	}

	for _, c := range cart {
		detailOrder := model.DetailPickupOnlineOrder{
			Id:                  uuid.New(),
			PickupOnlineOrderId: pickupOnlineOrder.Id,
			ProductId:           c.Id,
			Amount:              c.Amount,
			CreatedAt:           time.Now(),
			UpdatedAt:           time.Now(),
		}

		if err := s.pickupOnlineOrderRepository.SaveDetailPickupOnlineOrder(detailOrder); err != nil {
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

func (s *pickupOnlineOrderService) GetAllPickupOnlineOrder() ([]model.PickupOnlineOrder, error) {
	pickupOnlineOrders, err := s.pickupOnlineOrderRepository.GetAllPickupOnlineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return pickupOnlineOrders, err
	}

	return pickupOnlineOrders, nil
}

func (s *pickupOnlineOrderService) EditPickupOnlineOrder(id uuid.UUID) (model.EditPickupOnlineOrderResponse, error) {
	var editPickupOnlineOrderResponse model.EditPickupOnlineOrderResponse

	pickupOnlineOrder, err := s.pickupOnlineOrderRepository.GetPickupOnlineOrderById(id)
	if err != nil {
		log.Println("error: " + err.Error())
		return editPickupOnlineOrderResponse, err
	}

	detailPickupOnlineOrderResponse, err := s.pickupOnlineOrderRepository.GetDetailPickupOnlineOrderByPickupOnlineOrderId(id)
	if err != nil {
		log.Println("error: " + err.Error())
		return editPickupOnlineOrderResponse, err
	}

	for i, detailPickupOnlineOrder := range detailPickupOnlineOrderResponse {
		detailPickupOnlineOrderResponse[i].PriceFormatted = helper.FormatRupiah(float64(detailPickupOnlineOrder.Price))
	}

	editPickupOnlineOrderResponse.Id = id
	editPickupOnlineOrderResponse.Name = pickupOnlineOrder.Name
	editPickupOnlineOrderResponse.PhoneNumber = pickupOnlineOrder.PhoneNumber
	editPickupOnlineOrderResponse.Total = pickupOnlineOrder.Total
	editPickupOnlineOrderResponse.TotalFormatted = helper.FormatRupiah(float64(editPickupOnlineOrderResponse.Total))
	editPickupOnlineOrderResponse.Status = pickupOnlineOrder.Status
	editPickupOnlineOrderResponse.PickupDate = pickupOnlineOrder.PickupDate.Format("15:04:05 Monday, 02 January 2006")
	editPickupOnlineOrderResponse.DetailPickupOnlineOrderResponse = detailPickupOnlineOrderResponse

	return editPickupOnlineOrderResponse, nil
}

func (s *pickupOnlineOrderService) UpdateStatusPickupOnlineOrder(id uuid.UUID, status string) error {
	data := map[string]interface{}{
		"status":     status,
		"updated_at": time.Now(),
	}
	err := s.pickupOnlineOrderRepository.UpdatePickupOnlineOrder(id, data)
	if err != nil {
		log.Println("error: " + err.Error())
		return err
	}

	return nil
}

func (s *pickupOnlineOrderService) DeletePickupOnlineOrder(id uuid.UUID) error {
	pickupOnlineOrder, err := s.pickupOnlineOrderRepository.GetPickupOnlineOrderById(id)
	if err != nil {
		log.Println("error: " + err.Error())
		return err
	}

	detailPickupOnlineOrders, err := s.pickupOnlineOrderRepository.GetAllDetailPickupOnlineOrderByPickupOnlineOrderId(id)
	if err != nil {
		log.Println("error: " + err.Error())
		return err
	}

	for _, doo := range detailPickupOnlineOrders {
		if err := s.pickupOnlineOrderRepository.DeleteDetailPickupOnlineOrder(doo); err != nil {
			log.Println("error: " + err.Error())
			return err
		}
	}

	if err := s.pickupOnlineOrderRepository.DeletePickupOnlineOrder(pickupOnlineOrder); err != nil {
		log.Println("error: " + err.Error())
		return err
	}

	return nil
}

func (s *pickupOnlineOrderService) CountPickupOnlineOrderByStatus(status string) (int, error) {
	count, err := s.pickupOnlineOrderRepository.CountPickupOnlineOrderByStatus(status)
	if err != nil {
		log.Println("error: " + err.Error())
		return count, err
	}

	return count, nil
}
