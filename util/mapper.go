package util

import (
	"github.com/IsraelTeo/api-store-go/dto"
	"github.com/IsraelTeo/api-store-go/model"
)

func ToUserDTO(user *model.User) *dto.UserResponse {
	if user == nil {
		return nil
	}

	return &dto.UserResponse{
		Email:   user.Email,
		IsAdmin: user.IsAdmin,
	}
}

func ToListUserDTO(users []model.User) []dto.UserResponse {
	var userResponses []dto.UserResponse

	for _, user := range users {
		userResponses = append(userResponses, dto.UserResponse{
			Email:   user.Email,
			IsAdmin: user.IsAdmin,
		})
	}

	return userResponses
}
