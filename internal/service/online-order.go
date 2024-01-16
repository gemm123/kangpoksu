package service

import (
	"github.com/google/uuid"
	"kopoksu/helper"
	"kopoksu/internal/model"
	"kopoksu/internal/repository"
	"log"
	"time"
)

type onlineOrderService struct {
	onlineOrderRepository repository.OnlineOrderRepository
	productRepository     repository.ProductRepository
}

type OnlineOrderService interface {
	SaveOnlineOrder(onlineOrder model.OnlineOrder, cart []model.Cart) error
	GetAllOnlineOrder() ([]model.OnlineOrder, error)
	EditOnlineOrder(id uuid.UUID) (model.EditOnlineOrderResponse, error)
	UpdateStatusOnlineOrder(id uuid.UUID, status string) error
	DeleteOnlineOrder(id uuid.UUID) error
	CountOnlineOrderByStatus(status string) (int, error)
}

func NewOnlineOrderService(onlineOrderRepository repository.OnlineOrderRepository, productRepository repository.ProductRepository) *onlineOrderService {
	return &onlineOrderService{
		onlineOrderRepository: onlineOrderRepository,
		productRepository:     productRepository,
	}
}

func (s *onlineOrderService) SaveOnlineOrder(onlineOrder model.OnlineOrder, cart []model.Cart) error {
	onlineOrder.Id = uuid.New()
	onlineOrder.CreatedAt = time.Now()
	onlineOrder.UpdatedAt = time.Now()
	onlineOrder.Status = "Menunggu konfirmasi pembayaran"

	if err := s.onlineOrderRepository.SaveOnlineOrder(onlineOrder); err != nil {
		log.Println("error: " + err.Error())
		return err
	}

	for _, c := range cart {
		detailOrder := model.DetailOnlineOrder{
			Id:            uuid.New(),
			OnlineOrderId: onlineOrder.Id,
			ProductId:     c.Id,
			Amount:        c.Amount,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}

		if err := s.onlineOrderRepository.SaveDetailOnlineOrder(detailOrder); err != nil {
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

func (s *onlineOrderService) GetAllOnlineOrder() ([]model.OnlineOrder, error) {
	onlineOrders, err := s.onlineOrderRepository.GetAllOnlineOrder()
	if err != nil {
		log.Println("error: " + err.Error())
		return onlineOrders, err
	}

	return onlineOrders, nil
}

func (s *onlineOrderService) EditOnlineOrder(id uuid.UUID) (model.EditOnlineOrderResponse, error) {
	var editOnlineOrderResponse model.EditOnlineOrderResponse

	onlineOrder, err := s.onlineOrderRepository.GetOnlineOrderById(id)
	if err != nil {
		log.Println("error: " + err.Error())
		return editOnlineOrderResponse, err
	}

	detailOnlineOrderResponse, err := s.onlineOrderRepository.GetDetailOnlineOrderByOnlineOrderId(id)
	if err != nil {
		log.Println("error: " + err.Error())
		return editOnlineOrderResponse, err
	}

	for i, detailOfflineOrder := range detailOnlineOrderResponse {
		detailOnlineOrderResponse[i].PriceFormatted = helper.FormatRupiah(float64(detailOfflineOrder.Price))
	}

	editOnlineOrderResponse.Id = id
	editOnlineOrderResponse.Name = onlineOrder.Name
	editOnlineOrderResponse.Address = onlineOrder.Address
	editOnlineOrderResponse.City = onlineOrder.City
	editOnlineOrderResponse.Province = onlineOrder.Province
	editOnlineOrderResponse.PhoneNumber = onlineOrder.PhoneNumber
	editOnlineOrderResponse.PostCode = onlineOrder.PostCode
	editOnlineOrderResponse.Total = onlineOrder.Total
	editOnlineOrderResponse.TotalFormatted = helper.FormatRupiah(float64(editOnlineOrderResponse.Total))
	editOnlineOrderResponse.Status = onlineOrder.Status
	editOnlineOrderResponse.DetailOnlineOrderResponse = detailOnlineOrderResponse

	return editOnlineOrderResponse, nil
}

func (s *onlineOrderService) UpdateStatusOnlineOrder(id uuid.UUID, status string) error {
	data := map[string]interface{}{
		"status":     status,
		"updated_at": time.Now(),
	}
	err := s.onlineOrderRepository.UpdateOnlineOrder(id, data)
	if err != nil {
		log.Println("error: " + err.Error())
		return err
	}

	return nil
}

func (s *onlineOrderService) DeleteOnlineOrder(id uuid.UUID) error {
	onlineOrder, err := s.onlineOrderRepository.GetOnlineOrderById(id)
	if err != nil {
		log.Println("error: " + err.Error())
		return err
	}

	detailOnlineOrders, err := s.onlineOrderRepository.GetAllDetailOnlineOrderByOnlineOrderId(id)
	if err != nil {
		log.Println("error: " + err.Error())
		return err
	}

	for _, doo := range detailOnlineOrders {
		if err := s.onlineOrderRepository.DeleteDetailOnlineOrder(doo); err != nil {
			log.Println("error: " + err.Error())
			return err
		}
	}

	if err := s.onlineOrderRepository.DeleteOnlineOrder(onlineOrder); err != nil {
		log.Println("error: " + err.Error())
		return err
	}

	return nil
}

func (s *onlineOrderService) CountOnlineOrderByStatus(status string) (int, error) {
	count, err := s.onlineOrderRepository.CountOnlineOrderByStatus(status)
	if err != nil {
		log.Println("error: " + err.Error())
		return count, err
	}

	return count, nil
}
