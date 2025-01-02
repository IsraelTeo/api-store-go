package auth

import (
	"log"
	"net/http"

	"github.com/IsraelTeo/api-store-go/model"
	"github.com/IsraelTeo/api-store-go/repository"
	"github.com/IsraelTeo/api-store-go/response"
	"github.com/labstack/echo/v4"
)

type LoginService interface {
	Login(c echo.Context) error
}

type UserLogin struct {
	repo repository.UserRepository
}

func NewLogin(repo repository.UserRepository) LoginService {
	return &UserLogin{repo: repo}
}

func (l *UserLogin) Login(c echo.Context) error {
	credentials := model.RegisterUserPayload{}
	err := c.Bind(&credentials)
	if err != nil {
		log.Printf("Bad request: %v", err)
		return response.WriteError(c, http.StatusBadRequest, "Bad request", err)
	}

	user, err := l.repo.GetByEmail(credentials.Email)
	if err != nil {
		log.Printf("Invalid Email: %v", err)
		return response.WriteError(c, http.StatusUnauthorized, "Invalid email", err)
	}

	isValid, errorMessage := ComparePassword(user.Password, []byte(credentials.Password))
	if !isValid {
		log.Printf("Error: %v", errorMessage)
		return response.WriteError(c, http.StatusUnauthorized, errorMessage, nil)
	}

	token, err := GenerateToken(user)
	if err != nil {
		log.Printf("Error generating token: %v", err)
		return response.WriteError(c, http.StatusInternalServerError, "Error generating token", err)
	}

	responseMap := map[string]interface{}{
		"role":  user.IsAdmin,
		"token": token,
	}

	return response.WriteResponse(c, http.StatusOK, "Login successfully", responseMap)
}
