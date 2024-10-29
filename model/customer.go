package model

import (
	"gorm.io/gorm"
)

type Sales []Sale

type Customer struct {
	gorm.Model

	Name     string `json:"name"`
	LastName string `json:"last_name"`
	DNI      string `json:"dni"`
	Sales    Sales  `json:"sales" gorm:"foreignKey:CustomerID"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
