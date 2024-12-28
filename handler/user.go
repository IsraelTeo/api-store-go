package handler

import (
	"net/http"

	"github.com/IsraelTeo/api-store-go/model"
	"github.com/IsraelTeo/api-store-go/payload"
	"github.com/IsraelTeo/api-store-go/service"
	"github.com/IsraelTeo/api-store-go/validate"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUserByID(c echo.Context) error {
	ID, err := validate.ValidateAndParseID(c)
	if err != nil {
		response := payload.NewResponse(payload.MessageTypeError, "Invalid ID format.", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	user, err := h.service.GetBydID(uint(ID))
	if err != nil {
		response := payload.NewResponse(payload.MessageTypeError, "User not found", nil)
		return c.JSON(http.StatusNotFound, response)
	}

	response := payload.NewResponse(payload.MessageTypeSuccess, "User found successfuslly", user)
	return c.JSON(http.StatusOK, response)
}

func (h *UserHandler) GetAllUsers(c echo.Context) error {
	users, err := h.service.GetAll()
	if err != nil {
		response := payload.NewResponse(payload.MessageTypeError, "Failed to fetch users", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := payload.NewResponse(payload.MessageTypeSuccess, "Users found", users)
	return c.JSON(http.StatusOK, response)
}

func (h *UserHandler) RegisterUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		response := payload.NewResponse(payload.MessageTypeError, "Bad request", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := h.service.RegisterUser(&user); err != nil {
		response := payload.NewResponse(payload.MessageTypeError, "Failed to save user", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := payload.NewResponse(payload.MessageTypeSuccess, "User created successfully", nil)
	return c.JSON(http.StatusCreated, response)
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		response := payload.NewResponse(payload.MessageTypeError, "Bad request", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	ID, err := validate.ValidateAndParseID(c)
	if err != nil {
		response := payload.NewResponse(payload.MessageTypeError, "Invalid ID format", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	userUpdated, err := h.service.Update(uint(ID), &user)
	if err != nil {
		response := payload.NewResponse(payload.MessageTypeError, "Failed to update user", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := payload.NewResponse(payload.MessageTypeSuccess, "User updated successfully", userUpdated)
	return c.JSON(http.StatusOK, response)
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	ID, err := validate.ValidateAndParseID(c)
	if err != nil {
		response := payload.NewResponse(payload.MessageTypeError, "Invalid ID format", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := h.service.Delete(ID); err != nil {
		response := payload.NewResponse(payload.MessageTypeError, "Failed to delete user", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := payload.NewResponse(payload.MessageTypeSuccess, "User deleted successfully", nil)
	return c.JSON(http.StatusOK, response)
}
