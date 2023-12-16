package model

import (
	"github.com/google/uuid"
	"time"
)

type DetailOrder struct {
	Id        uuid.UUID `gorm:"column:"`
	OrderId   uuid.UUID `gorm:"column:order_id"`
	ProductId uuid.UUID `gorm:"column:product_id"`
	Amount    int       `gorm:"column:amount"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
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

type OnlineOrder struct {
	Id          uuid.UUID `gorm:"column:id"`
	Name        string    `gorm:"column:name" form:"name"`
	Address     string    `gorm:"column:address" form:"address"`
	City        string    `gorm:"column:city" form:"city"`
	Province    string    `gorm:"column:province" form:"province"`
	PhoneNumber string    `gorm:"column:phone_number" form:'number'`
	PostCode    int       `gorm:"column:post_code" form:"post-code"`
	Total       int       `gorm:"column:total"`
	Cost        int       `gorm:"column:cost"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}
