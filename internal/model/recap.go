package model

type RecapSalesByMonth struct {
	Month int `gorm:"column:bulan"`
	Sold  int `gorm:"column:terjual"`
}
