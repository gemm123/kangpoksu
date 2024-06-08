package model

import (
	"github.com/google/uuid"
	"time"
)

type DetailPickupOnlineOrder struct {
	Id                  uuid.UUID `gorm:"column:id"`
	PickupOnlineOrderId uuid.UUID `gorm:"column:pickup_online_order_id"`
	ProductId           uuid.UUID `gorm:"column:product_id"`
	Amount              int       `gorm:"column:amount"`
	CreatedAt           time.Time `gorm:"column:created_at"`
	UpdatedAt           time.Time `gorm:"column:updated_at"`
}

type DetailPickupOnlineOrderResponse struct {
	Name           string
	Amount         int
	Price          int
	PriceFormatted string
}

type PickupOnlineOrder struct {
	Id            uuid.UUID `gorm:"column:id"`
	Name          string    `gorm:"column:name" form:"name"`
	PhoneNumber   string    `gorm:"column:phone_number" form:"number"`
	Total         int       `gorm:"column:total"`
	Status        string    `gorm:"column:status"`
	PickupDateStr string    `gorm:"-" form:"pickup-date"`
	PickupDate    time.Time `gorm:"column:pickup_date"`
	CreatedAt     time.Time `gorm:"column:created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at"`
}

type EditPickupOnlineOrderResponse struct {
	Id                              uuid.UUID
	Name                            string
	PhoneNumber                     string
	Total                           int
	TotalFormatted                  string
	Status                          string
	PickupDate                      string
	DetailPickupOnlineOrderResponse []DetailPickupOnlineOrderResponse
}
