package model

import (
	"github.com/google/uuid"
	"time"
)

type DetailOfflineOrder struct {
	Id             uuid.UUID `gorm:"column:id"`
	OfflineOrderId uuid.UUID `gorm:"column:offline_order_id"`
	ProductId      uuid.UUID `gorm:"column:product_id"`
	Amount         int       `gorm:"column:amount"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
}

type DetailOfflineOrderResponse struct {
	Name   string
	Amount int
	Price  int
}

type OfflineOrder struct {
	Id          uuid.UUID `gorm:"column:id"`
	Name        string    `gorm:"column:name" form:"name"`
	PhoneNumber string    `gorm:"column:phone_number" form:"number"`
	Total       int       `gorm:"column:total"`
	Status      string    `gorm:"column:status"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

type EditOfflineOrderResponse struct {
	Id                         uuid.UUID
	Name                       string
	PhoneNumber                string
	Total                      int
	Status                     string
	DetailOfflineOrderResponse []DetailOfflineOrderResponse
}
