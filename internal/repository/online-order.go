package repository

import (
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
