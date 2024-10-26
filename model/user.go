package model

import "gorm.io/gorm"

type Role string

var (
	AdminRole    Role = "admin"
	CustomerRole Role = "customer"
)

type User struct {
	gorm.Model

	Username string   `json:"username"`
	Password string   `json:"password"`
	Roles    []Role   `json:"roles"`
	Customer Customer `gorm:"foreingKey:UserID"`
}
