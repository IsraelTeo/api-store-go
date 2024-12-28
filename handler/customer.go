package handler

import (
	"log"
	"net/http"

	"github.com/IsraelTeo/api-store-go/model"
	"github.com/IsraelTeo/api-store-go/payload"
	"github.com/IsraelTeo/api-store-go/service"
	"github.com/IsraelTeo/api-store-go/validate"
	"github.com/labstack/echo/v4"
)

type CustomerHandler struct {
	customerService service.CustomerService
}

func NewCustomerHandler(customerService service.CustomerService) *CustomerHandler {
	return &CustomerHandler{customerService: customerService}
}

func (h *CustomerHandler) GetCustomerByID(c echo.Context) error {
	ID, err := validate.ValidateAndParseID(c)
	if err != nil {
		response := payload.NewResponse(payload.MessageTypeError, "Invalid ID format.", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	customer, err := h.customerService.GetByID(uint(ID))
	if err != nil {
		response := payload.NewResponse(payload.MessageTypeError, "Customer not found", nil)
		return c.JSON(http.StatusNotFound, response)
	}

	response := payload.NewResponse(payload.MessageTypeSuccess, "Customer found", customer)
	return c.JSON(http.StatusOK, response)
}

func (h *CustomerHandler) GetAllCustomers(c echo.Context) error {
	customers, err := h.customerService.GetAll()
	if err != nil {
		response := payload.NewResponse(payload.MessageTypeError, "Failed to fetch customers", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := payload.NewResponse(payload.MessageTypeSuccess, "Customers found", customers)
	return c.JSON(http.StatusOK, response)
}

func (h *CustomerHandler) CreateCustomer(c echo.Context) error {
	customer := model.Customer{}
	if err := c.Bind(&customer); err != nil {
		log.Printf("Error decoding request body: %v", err)
		response := payload.NewResponse(payload.MessageTypeError, "Bad request", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := h.customerService.Create(&customer); err != nil {
		response := payload.NewResponse(payload.MessageTypeError, "Failed to save customer", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := payload.NewResponse(payload.MessageTypeSuccess, "Customer created successfully", nil)
	return c.JSON(http.StatusCreated, response)
}

func (h *CustomerHandler) UpdateCustomer(c echo.Context) error {
	ID, err := validate.ValidateAndParseID(c)
	if err != nil {
		response := payload.NewResponse(payload.MessageTypeError, "Invalid ID format", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	customer := model.Customer{}
	if err := c.Bind(&customer); err != nil {
		log.Printf("Error decoding request body: %v", err)
		response := payload.NewResponse(payload.MessageTypeError, "Bad request", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	updatedCustomer, err := h.customerService.Update(uint(ID), &customer)
	if err != nil {
		log.Printf("Error updating customer: %v", err)
		response := payload.NewResponse(payload.MessageTypeError, "Failed to update customer", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := payload.NewResponse(payload.MessageTypeSuccess, "Customer updated successfully", updatedCustomer)
	return c.JSON(http.StatusOK, response)
}

func (h *CustomerHandler) DeleteCustomer(c echo.Context) error {
	ID, err := validate.ValidateAndParseID(c)
	if err != nil {
		response := payload.NewResponse(payload.MessageTypeError, "Invalid ID format", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := h.customerService.Delete(uint(ID)); err != nil {
		response := payload.NewResponse(payload.MessageTypeError, "Failed to delete customer", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := payload.NewResponse(payload.MessageTypeSuccess, "Customer deleted successfully", nil)
	return c.JSON(http.StatusOK, response)
}
