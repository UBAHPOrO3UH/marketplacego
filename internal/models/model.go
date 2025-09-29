package models

//import (
//	"gorm.io/gorm"
//)

type Product struct {
	ID                uint    `gorm:"primaryKey" json:"id"`
	Title             string  `json:"title"`
	Description       string  `json:"description"`
	Price             float64 `json:"price"`
	InStock           bool    `json:"in_stock"`
	Discount          float64 `json:"discount"`
	PriceWithDiscount float64 `json:"price_with_discount"`
}
