package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ComparePassword(passwordHashed string, password []byte) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHashed), password)
	if err != nil {
		return false, err.Error()
	}
	return true, ""
}
