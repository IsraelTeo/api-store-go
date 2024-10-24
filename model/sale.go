package model

import "gorm.io/gorm"

type Sale struct {
	gorm.Model

	TotalAmount float64   `json:"total_amount"`
	Products    []Product `json:"products" gorm:"many2many:sale_products;"`
	CustomerID  uint      `json:"customer_id"`
	Customer    Customer  `json:"customer"`
}
