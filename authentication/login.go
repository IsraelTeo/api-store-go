package authentication

import (
	"log"
	"net/http"

	"github.com/IsraelTeo/api-store-go/payload"
	"github.com/IsraelTeo/api-store-go/repository"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type LoginService interface {
	Login(c echo.Context) error
}

type UserLogin struct {
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) LoginService {
	return &UserLogin{repo: repo}
}

func (l *UserLogin) Login(c echo.Context) error {
	var credentials Credentials

	err := c.Bind(&credentials)
	if err != nil {
		log.Printf("Bad request: %v", err)
		response := payload.NewResponse(payload.MessageTypeError, "Bad request", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	user, err := l.repo.GetByEmail(credentials.Email)
	if err != nil {
		log.Printf("Invalid Email: %v", err)
		response := payload.NewResponse(payload.MessageTypeError, "Invalid Email", nil)
		return c.JSON(http.StatusUnauthorized, response)
	}

	err = verifyPassword(user.Password, credentials.Password)
	if err != nil {
		log.Printf("Invalid password: %v", err)
		response := payload.NewResponse(payload.MessageTypeError, "Invalid password", nil)
		return c.JSON(http.StatusUnauthorized, response)
	}

	token, err := GenerateToken(*user)
	if err != nil {
		log.Printf("Error generating token: %v", err)
		response := payload.NewResponse(payload.MessageTypeError, "Error generating token", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	responseMap := map[string]interface{}{
		"role":  user.IsAdmin,
		"token": token,
	}

	response := payload.NewResponse(payload.MessageTypeSuccess, "Login successfully", responseMap)
	return c.JSON(http.StatusOK, response)
}

func verifyPassword(passwordHashed string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHashed), []byte(password))
	if err != nil {
		return err
	}

	return nil
}
