package util

import (
	"github.com/labstack/echo/v4"
)

/*type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponse(message string, data interface{}) Response {
	return Response{
		Message: "message",
		Data:    data,
	}
}*/

func WriteResponse(c echo.Context, status int, message string, data interface{}) error {
	return c.JSON(status, map[string]interface{}{
		"message": message,
		"data":    data,
	})
}

func WriteError(c echo.Context, status int, message string, err error) error {
	return c.JSON(status, map[string]string{
		"error":   message,
		"details": err.Error(),
	})
}
