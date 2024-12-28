package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"size:100;unique;not_null" validate:"required,email,max=100"`
	Password string `json:"password" gorm:"size:100" validate:"required,min=8,max=100"`
	IsAdmin  bool   `json:"is_admin" gorm:"default:false"`
}
