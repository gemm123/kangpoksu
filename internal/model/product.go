package model

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	Id          uuid.UUID `gorm:"column:id"`
	CategoryId  uuid.UUID `gorm:"column:category_id"`
	Name        string    `gorm:"column:name"`
	Description string    `gorm:"column:description"`
	Quantity    int       `gorm:"column:quantity"`
	Price       int       `gorm:"column:price"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}
