package model

import "gorm.io/gorm"

type Sale struct {
	gorm.Model
	TotalAmount float64
	Products    []Product `gorm:"many2many:sale_products;"`
	CustomerID  uint
	Customer    Customer
}
