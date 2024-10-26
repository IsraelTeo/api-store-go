package authorization

import (
	"github.com/IsraelTeo/api-store-go/model"
	"golang.org/x/crypto/bcrypt"
)

func Register(username, password string, roles []model.Role) (*model.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Username: username,
		Password: string(hashedPassword),
		Roles:    roles,
	}

	return user, nil
}
