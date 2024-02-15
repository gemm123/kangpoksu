package model

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	Id                uuid.UUID `gorm:"column:id"`
	CategoryId        uuid.UUID `gorm:"column:category_id"`
	Name              string    `gorm:"column:name" form:"product-name"`
	Description       string    `gorm:"column:description" form:"description"`
	Quantity          int       `gorm:"column:quantity" form:"quantity"`
	Weight            int       `gorm:"column:weight" form:"weight"`
	Price             int       `gorm:"column:price" form:"price"`
	BuyPrice          int       `gorm:"column:buy_price" form:"buy-price"`
	Image             string    `gorm:"column:image"`
	CreatedAt         time.Time `gorm:"column:created_at"`
	UpdatedAt         time.Time `gorm:"column:updated_at"`
	PriceFormatted    string    `gorm:"-"`
	BuyPriceFormatted string    `gorm:"-"`
}

type Cart struct {
	Id             uuid.UUID `form:"id" json:"id"`
	Name           string    `json:"name"`
	Amount         int       `form:"amount" json:"amount"`
	Total          int       `json:"total"`
	Image          string    `json:"image"`
	PriceFormatted string
}

type SearchResult struct {
	Id       uuid.UUID `json:"id" gorm:"id"`
	Name     string    `json:"name" gorm:"name"`
	Category string    `json:"category" gorm:"category"`
}
