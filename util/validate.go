package util

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/IsraelTeo/api-store-go/db"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// CustomValidator estructura para validación
type CustomValidator struct {
	Validator *validator.Validate
}

// Validate implementa el método de la interfaz echo.Validator
func (v *CustomValidator) Validate(i interface{}) error {
	return v.Validator.Struct(i)
}

// InitValidator inicializa el validador
func InitValidator() *CustomValidator {
	return &CustomValidator{
		Validator: validator.New(),
	}
}

func VerifyListEmpty[T any](list []T) bool {
	return len(list) == 0
}

func checkIfFieldExists[T any](field string, value interface{}, model *T) (bool, error) {
	err := db.GDB.Where(field+" = ?", value).First(model).Error
	if err == nil {
		return true, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	return false, err
}

func ValidateUniqueField[T any](field, value string, model *T) (bool, error) {
	exists, err := checkIfFieldExists(field, value, model)
	return exists, err
}

func IsEmpty(s string) bool {
	return s == ""
}

func ParseID(c echo.Context) (uint, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		return 0, fmt.Errorf("the id must not be less than 1 and must be numeric")
	}
	return uint(id), nil
}
