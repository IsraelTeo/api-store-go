package model

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model

	Name     string `json:"name"`
	LastName string `json:"last_name"`
	DNI      string `json:"dni"`
	Sales    []Sale `json:"sales" gorm:"foreignKey:CustomerID"`
}
