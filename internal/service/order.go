package service

import (
	"github.com/google/uuid"
	"kopoksu/internal/model"
	"kopoksu/internal/repository"
	"log"
	"time"
)

type orderService struct {
	orderRepository repository.OrderRepository
}

type OrderService interface {
	SaveOfflineOrder(offlineOrder model.OfflineOrder, cart []model.Cart) error
}

func NewOrderService(orderRepository repository.OrderRepository) *orderService {
	return &orderService{
		orderRepository: orderRepository,
	}
}

func (s *orderService) SaveOfflineOrder(offlineOrder model.OfflineOrder, cart []model.Cart) error {
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
	}

	//reduce quantity products
	//delete session cart

	return nil
}
