package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	ImageURL    string  `json:"image_url"`
}

type Inventory struct {
	gorm.Model
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Location  string  `json:"location"`
	Product   Product `gorm:"foreignKey:ProductID;references:ID"`
}

type Order struct {
	gorm.Model
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	OrderDate string  `json:"order_date"`
	Product   Product `gorm:"foreignKey:ProductID;references:ID"`
}
