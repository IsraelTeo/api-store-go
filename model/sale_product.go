package model

type SaleProduct struct {
	SaleID    uint `gorm:"primaryKey"`
	ProductID uint `gorm:"primaryKey"`
	Quantity  int  `json:"quantity" validate:"required,gt=0"`
}
