package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" gorm:"size:100;unique;not_null" validate:"required,email,max=100"`
	Password  string `json:"password" gorm:"size:100" validate:"required,min=8,max=100"`
	IsAdmin   bool   `json:"is_admin" gorm:"default:false"`
}

type RegisterUserPayload struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required, email"`
	Password  string `json:"password" validate:"required"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required, email"`
	Password string `json:"password" validate:"required"`
}
