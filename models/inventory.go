package models

type Inventory struct {
	ID        uint   `gorm:"primaryKey"`
	ProductID uint   `gorm:"not null"`
	Quantity  int    `gorm:"not null"`
	Location  string `gorm:"not null"`
}

func (Inventory) TableName() string {
	return "inventories"
}
