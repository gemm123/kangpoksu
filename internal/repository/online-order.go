package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"kopoksu/internal/model"
)

type onlineOrderRepository struct {
	DB *gorm.DB
}

type OnlineOrderRepository interface {
	SaveOnlineOrder(onlineOrder model.OnlineOrder) error
	SaveDetailOnlineOrder(detailOnlineOrder model.DetailOnlineOrder) error
	GetAllOnlineOrder() ([]model.OnlineOrder, error)
	GetOnlineOrderById(id uuid.UUID) (model.OnlineOrder, error)
	GetDetailOnlineOrderByOnlineOrderId(id uuid.UUID) ([]model.DetailOnlineOrderResponse, error)
}

func NewOnlineOrderRepository(DB *gorm.DB) *onlineOrderRepository {
	return &onlineOrderRepository{
		DB: DB,
	}
}

func (r *onlineOrderRepository) SaveOnlineOrder(onlineOrder model.OnlineOrder) error {
	err := r.DB.Create(&onlineOrder).Error
	return err
}

func (r *onlineOrderRepository) SaveDetailOnlineOrder(detailOnlineOrder model.DetailOnlineOrder) error {
	err := r.DB.Create(&detailOnlineOrder).Error
	return err
}

func (r *onlineOrderRepository) GetAllOnlineOrder() ([]model.OnlineOrder, error) {
	var onlineOrders []model.OnlineOrder
	err := r.DB.Find(&onlineOrders).Error
	return onlineOrders, err
}

func (r *onlineOrderRepository) GetOnlineOrderById(id uuid.UUID) (model.OnlineOrder, error) {
	var onlineOrder model.OnlineOrder
	err := r.DB.First(&onlineOrder, "id = ?", id).Error
	return onlineOrder, err
}

func (r *onlineOrderRepository) GetDetailOnlineOrderByOnlineOrderId(id uuid.UUID) ([]model.DetailOnlineOrderResponse, error) {
	var detailOnlineOrdersResponse []model.DetailOnlineOrderResponse
	err := r.DB.Table("detail_online_orders").
		Select("products.name, detail_online_orders.amount, detail_online_orders.amount * products.price as price").
		Joins("LEFT JOIN products ON products.id = detail_online_orders.product_id").
		Where("detail_online_orders.online_order_id = ?", id).
		Scan(&detailOnlineOrdersResponse).
		Error
	return detailOnlineOrdersResponse, err
}
