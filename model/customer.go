package model

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name     string
	LastName string
	DNI      string
	Sales    []Sale `gorm:"foreignKey:CustomerID"`
}
