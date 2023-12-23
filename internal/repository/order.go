package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"kopoksu/internal/model"
)

type orderRepository struct {
	DB *gorm.DB
}

type OrderRepository interface {
	SaveOfflineOrder(offlineOrder model.OfflineOrder) error
	SaveDetailOfflineOrder(detailOfflineOrder model.DetailOfflineOrder) error
	GetAllOfflineOrder() ([]model.OfflineOrder, error)
	GetOfflineOrderById(id uuid.UUID) (model.OfflineOrder, error)
	GetDetailOfflineOrderByOfflineOrderId(id uuid.UUID) ([]model.DetailOfflineOrderResponse, error)
}

func NewOrderRepository(DB *gorm.DB) *orderRepository {
	return &orderRepository{
		DB: DB,
	}
}

func (r *orderRepository) SaveOfflineOrder(offlineOrder model.OfflineOrder) error {
	err := r.DB.Create(&offlineOrder).Error
	return err
}

func (r *orderRepository) SaveDetailOfflineOrder(detailOfflineOrder model.DetailOfflineOrder) error {
	err := r.DB.Create(&detailOfflineOrder).Error
	return err
}

func (r *orderRepository) GetAllOfflineOrder() ([]model.OfflineOrder, error) {
	var offlineOrders []model.OfflineOrder
	err := r.DB.Find(&offlineOrders).Error
	return offlineOrders, err
}

func (r *orderRepository) GetOfflineOrderById(id uuid.UUID) (model.OfflineOrder, error) {
	var offlineOrder model.OfflineOrder
	err := r.DB.First(&offlineOrder, "id = ?", id).Error
	return offlineOrder, err
}

func (r *orderRepository) GetDetailOfflineOrderByOfflineOrderId(id uuid.UUID) ([]model.DetailOfflineOrderResponse, error) {
	var detailOfflineOrdersResponse []model.DetailOfflineOrderResponse
	err := r.DB.Table("detail_offline_orders").
		Select("products.name, detail_offline_orders.amount, detail_offline_orders.amount * products.price as price").
		Joins("LEFT JOIN products ON products.id = detail_offline_orders.product_id").
		Where("detail_offline_orders.offline_order_id = ?", id).
		Scan(&detailOfflineOrdersResponse).
		Error
	return detailOfflineOrdersResponse, err
}
