package auth

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/IsraelTeo/api-store-go/model"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GenerateToken(user *model.User) (string, error) {
	expiration := os.Getenv("JWT_EXP")
	secret := os.Getenv("API_SECRET")
	issuedAt := time.Now().Unix()

	if expiration == "" || secret == "" {
		return "", fmt.Errorf("missing JWT_EXP or API_SECRET environment variable")
	}

	// Crear payload con claims
	payload := jwt.MapClaims{
		"email":      user.Email,   // Correo del usuario
		"authorized": true,         // Indica si el usuario está autorizado
		"is_admin":   user.IsAdmin, // Indica si el usuario es administrador
		"iat":        issuedAt,     // Tiempo actual en formato Unix (emisión del token)
		"exp":        expiration,   // Expiración del token
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload) // Crea un nuevo token usando el algoritmo de firma HS256 y el payload
	tokenString, err := token.SignedString([]byte(secret))      // Firma el token con una clave secreta obtenida de las variables de entorno
	if err != nil {
		log.Printf("Error signing the token: %\n", err)
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(c echo.Context) (model.User, error) {
	token, err := GetToken(c)
	if err != nil {
		log.Printf("Error retrieving token: %v", err)
		return model.User{}, fmt.Errorf("no token found in request: %w", err)
	}

	jwtToken, err := jwt.Parse(token, validateMethodAndGetSecret) //verifica que el token sea válido
	if err != nil {
		log.Printf("Token not valid: %v\n", err)
		return model.User{}, fmt.Errorf("invalid token: %w", err)
	}

	userData, ok := jwtToken.Claims.(jwt.MapClaims) //verificamos que los claims sean del tipo jwt.MapClaims
	if !ok || !jwtToken.Valid {
		log.Println("Unable to retrieve payload information or token is invalid")
		return model.User{}, fmt.Errorf("invalid token claims")
	}

	_, ok = userData["email"].(string) //verificamos que el email sea string
	if !ok {
		log.Println("Email field missing or not a string in token claims")
		return model.User{}, fmt.Errorf("email field is missing or invalid in token claims")
	}

	response := model.User{
		Email:   userData["email"].(string),  //asignamos el email y nos aseguramos que sea un string
		IsAdmin: userData["is_admin"].(bool), //asignamos el rol admin y nos aseguramos que sea un bool
	}

	return response, nil
}

func GetToken(c echo.Context) (string, error) {
	token := c.QueryParam("token")
	if token != "" {
		return token, nil
	}

	authHeader := c.Request().Header.Get("Authorization")
	if len(strings.Split(authHeader, " ")) == 2 {
		return strings.Split(authHeader, " ")[1], nil
	}

	return "", errors.New("token not found in request")
}

func validateMethodAndGetSecret(token *jwt.Token) (any, error) {
	_, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return nil, fmt.Errorf("method not valid")
	}

	return []byte(os.Getenv("API_SECRET")), nil
}
