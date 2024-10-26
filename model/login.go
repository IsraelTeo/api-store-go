package model

import (
	"github.com/golang-jwt/jwt/v5"
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claim struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}
