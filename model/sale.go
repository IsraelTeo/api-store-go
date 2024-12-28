package model

import "gorm.io/gorm"

type Sale struct {
	gorm.Model
	TotalAmount float64   `json:"-" validate:"required,gt=0"`
	Products    []Product `gorm:"many2many:sale_products;" json:"products" validate:"required,dive"`
	CustomerID  uint      `json:"id_customer" gorm:"foreignKey:CustomerID" validate:"required"`
	Customer    Customer  `gorm:"foreignKey:CustomerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"customer"`
}
