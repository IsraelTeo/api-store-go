package handler

import (
	"net/http"

	"github.com/IsraelTeo/api-store-go/model"
	"github.com/IsraelTeo/api-store-go/service"
	"github.com/IsraelTeo/api-store-go/util"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUserByID(c echo.Context) error {
	ID, err := util.ParseID(c)
	if err != nil {
		return util.WriteError(c, http.StatusBadRequest, "Invalid ID.", nil)
	}

	user, err := h.service.GetByID(uint(ID))
	if err != nil {
		return util.WriteError(c, http.StatusNotFound, "User not found", nil)
	}

	return util.WriteResponse(c, http.StatusOK, "User found successfully", user)
}

func (h *UserHandler) GetAllUsers(c echo.Context) error {
	users, err := h.service.GetAll()
	if err != nil {
		return util.WriteError(c, http.StatusInternalServerError, "Failed to fetch users", err)
	}

	return util.WriteResponse(c, http.StatusOK, "Users found", users)
}

func (h *UserHandler) RegisterUser(c echo.Context) error {
	user := model.RegisterUserPayload{}
	if err := c.Bind(&user); err != nil {
		return util.WriteError(c, http.StatusBadRequest, "Bad request", err)
	}

	if err := h.service.RegisterUser(&user); err != nil {
		return util.WriteError(c, http.StatusInternalServerError, "Failed to save user", err)
	}

	return util.WriteResponse(c, http.StatusCreated, "User created successfully", nil)
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	ID, err := util.ParseID(c)
	if err != nil {
		return util.WriteError(c, http.StatusBadRequest, "Invalid ID format", err)
	}

	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return util.WriteError(c, http.StatusBadRequest, "Bad request", err)
	}

	userUpdated, err := h.service.Update(uint(ID), &user)
	if err != nil {
		return util.WriteError(c, http.StatusInternalServerError, "Failed to update user", err)
	}

	return util.WriteResponse(c, http.StatusOK, "User updated successfully", userUpdated)
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	ID, err := util.ParseID(c)
	if err != nil {
		return util.WriteError(c, http.StatusBadRequest, "Invalid ID format", err)
	}

	if err := h.service.Delete(ID); err != nil {
		return util.WriteError(c, http.StatusInternalServerError, "Failed to delete user", err)
	}

	return util.WriteResponse(c, http.StatusOK, "User deleted successfully", nil)
}
