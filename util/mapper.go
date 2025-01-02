package util

import (
	"github.com/IsraelTeo/api-store-go/model"
)

func ToUser(user *model.RegisterUserPayload) *model.User {
	if user == nil {
		return nil
	}

	return &model.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
		IsAdmin:   user.IsAdmin,
	}
}

func ToUserUpdated(user *model.User) *model.RegisterUserPayload {
	if user == nil {
		return nil
	}

	return &model.RegisterUserPayload{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
		IsAdmin:   user.IsAdmin,
	}
}
