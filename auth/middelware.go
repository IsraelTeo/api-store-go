package auth

import (
	"log"
	"net/http"

	"github.com/IsraelTeo/api-store-go/response"
	"github.com/labstack/echo/v4"
)

func ValidateJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, err := ValidateToken(c)
		if err != nil {
			log.Printf("Invalid token: %v", err)
			return response.WriteError(c, http.StatusUnauthorized, "Invalid token.", nil)
		}
		return next(c)
	}
}

func ValidateJWTAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userData, err := ValidateToken(c)
		if err != nil {
			log.Printf("Invalid token %v", err)
			return response.WriteError(c, http.StatusUnauthorized, "Invalid token.", nil)
		}

		if !userData.IsAdmin {
			log.Printf("Not admin %v", err)
			return response.WriteError(c, http.StatusForbidden, "Not admin", nil)
		}

		return next(c)
	}
}
