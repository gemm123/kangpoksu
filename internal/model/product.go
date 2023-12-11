package model

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	Id          uuid.UUID `gorm:"column:id"`
	CategoryId  uuid.UUID `gorm:"column:category_id"`
	Name        string    `gorm:"column:name" form:"product-name"`
	Description string    `gorm:"column:description" form:"description"`
	Quantity    int       `gorm:"column:quantity" form:"quantity"`
	Price       int       `gorm:"column:price" form:"price"`
	Image       string    `gorm:"column:image"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

type Cart struct {
	Id     uuid.UUID `form:"id" json:"id"`
	Name   string    `json:"name"`
	Amount int       `form:"amount" json:"amount"`
	Total  int       `json:"total"`
	Image  string    `json:"image"`
}
