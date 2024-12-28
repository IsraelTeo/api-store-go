package model

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model

	Name     string `json:"name" validate:"required,min=2,max=50"`
	LastName string `json:"last_name" validate:"required,min=2,max=70"`
	DNI      string `json:"dni" gorm:"unique" validate:"required,max=15,numeric"`
	Sales    []Sale `gorm:"foreignKey:CustomerID" json:"sales"`
}
