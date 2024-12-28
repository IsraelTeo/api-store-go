package authentication

import (
	"log"
	"net/http"

	"github.com/IsraelTeo/api-store-go/payload"
	"github.com/labstack/echo/v4"
)

func ValidateJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, err := ValidateToken(c)
		if err != nil {
			log.Printf("Invalid token: %v", err)
			response := payload.NewResponse(payload.MessageTypeError, "Invalid token.", nil)
			return c.JSON(http.StatusUnauthorized, response)
		}
		return next(c)
	}
}

func ValidateJWTAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userData, err := ValidateToken(c)
		if err != nil {
			log.Printf("Invalid token %v", err)
			response := payload.NewResponse(payload.MessageTypeError, "Invalid token", nil)
			return c.JSON(http.StatusUnauthorized, response)
		}

		if !userData.IsAdmin {
			log.Printf("Not admin %v", err)
			response := payload.NewResponse(payload.MessageTypeError, "Not admin", nil)
			return c.JSON(http.StatusForbidden, response)
		}

		return next(c)
	}
}
