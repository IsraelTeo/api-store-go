package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" gorm:"size:100;unique;not_null" validate:"required,email,max=100"`
	Password  string `json:"-" gorm:"size:100"`
	IsAdmin   bool   `json:"is_admin" gorm:"default:false"`
}

type RegisterUserPayload struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email,min=8,max=100"`
	Password  string `json:"password" validate:"required"`
	IsAdmin   bool   `json:"is_admin"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required, email"`
	Password string `json:"password" validate:"required"`
}
