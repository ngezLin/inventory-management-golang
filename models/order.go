package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ProductID uint      `json:"product_id"`
	Quantity  int       `json:"quantity"`
	OrderDate time.Time `json:"order_date"`
	Product   Product   `gorm:"foreignKey:ProductID"`
}
