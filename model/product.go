package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	Name  string  `json:"name"`
	Mark  string  `json:"mark"`
	Price float64 `json:"price"`
}
