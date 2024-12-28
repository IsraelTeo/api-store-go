package model

import "gorm.io/gorm"

type Stock []Product

type Product struct {
	gorm.Model
	Name  string  `json:"name" validate:"required,min=2,max=50"`
	Mark  string  `json:"mark" validate:"required"`
	Price float64 `json:"price" validate:"required,gt=0"`
}
