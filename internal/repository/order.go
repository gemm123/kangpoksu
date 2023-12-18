package repository

import (
	"gorm.io/gorm"
	"kopoksu/internal/model"
)

type orderRepository struct {
	DB *gorm.DB
}

type OrderRepository interface {
	SaveOfflineOrder(offlineOrder model.OfflineOrder) error
	SaveDetailOfflineOrder(detailOfflineOrder model.DetailOfflineOrder) error
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
