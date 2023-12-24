package model

import (
	"github.com/google/uuid"
	"time"
)

type OnlineOrder struct {
	Id          uuid.UUID `gorm:"column:id"`
	Name        string    `gorm:"column:name" form:"name"`
	Address     string    `gorm:"column:address" form:"address"`
	City        string    `gorm:"column:city" form:"city"`
	Province    string    `gorm:"column:province" form:"province"`
	PhoneNumber string    `gorm:"column:phone_number" form:"phone-number"`
	PostCode    int       `gorm:"column:post_code" form:"post-code"`
	Total       int       `gorm:"column:total"`
	Cost        int       `gorm:"column:cost"`
	Status      string    `gorm:"column:status"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

type DetailOnlineOrder struct {
	Id            uuid.UUID `gorm:"column:id"`
	OnlineOrderId uuid.UUID `gorm:"column:online_order_id"`
	ProductId     uuid.UUID `gorm:"column:product_id"`
	Amount        int       `gorm:"column:amount"`
	CreatedAt     time.Time `gorm:"column:created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at"`
}

type DetailOnlineOrderResponse struct {
	Name   string
	Amount int
	Price  int
}

type EditOnlineOrderResponse struct {
	Id                        uuid.UUID
	Name                      string
	Address                   string
	City                      string
	Province                  string
	PhoneNumber               string
	PostCode                  int
	Total                     int
	Status                    string
	DetailOnlineOrderResponse []DetailOnlineOrderResponse
}
