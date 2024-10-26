package authorization

import (
	"errors"

	"github.com/IsraelTeo/api-store-go/model"
	"golang.org/x/crypto/bcrypt"
)

func Login(username, password string) (string, error) {
	var user model.Login
	if user.Email == "" {
		return "", errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid password")
	}

	return GenerateToken(&user)
}
