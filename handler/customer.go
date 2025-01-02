package handler

import (
	"log"
	"net/http"

	"github.com/IsraelTeo/api-store-go/model"
	"github.com/IsraelTeo/api-store-go/service"
	"github.com/IsraelTeo/api-store-go/util"
	"github.com/labstack/echo/v4"
)

type CustomerHandler struct {
	customerService service.CustomerService
}

func NewCustomerHandler(customerService service.CustomerService) *CustomerHandler {
	return &CustomerHandler{customerService: customerService}
}

func (h *CustomerHandler) GetCustomerByID(c echo.Context) error {
	ID, err := util.ParseID(c)
	if err != nil {
		return util.WriteError(c, http.StatusBadRequest, "Invalid ID", nil)
	}

	customer, err := h.customerService.GetByID(uint(ID))
	if err != nil {
		return util.WriteError(c, http.StatusNotFound, "Customer not found", nil)
	}

	return util.WriteResponse(c, http.StatusOK, "Customer found", customer)
}

func (h *CustomerHandler) GetAllCustomers(c echo.Context) error {
	customers, err := h.customerService.GetAll()
	if err != nil {
		return util.WriteError(c, http.StatusInternalServerError, "Failed to fetch customers", err)
	}

	return util.WriteResponse(c, http.StatusOK, "Customers found", customers)
}

func (h *CustomerHandler) CreateCustomer(c echo.Context) error {
	customer := model.Customer{}
	if err := c.Bind(&customer); err != nil {
		return util.WriteError(c, http.StatusBadRequest, "Bad request", err)
	}

	if err := h.customerService.Create(&customer); err != nil {
		return util.WriteError(c, http.StatusInternalServerError, "Failed to save customer", err)
	}

	return util.WriteResponse(c, http.StatusCreated, "Customer created successfully", nil)
}

func (h *CustomerHandler) UpdateCustomer(c echo.Context) error {
	ID, err := util.ParseID(c)
	if err != nil {
		return util.WriteError(c, http.StatusBadRequest, "Invalid ID format", err)
	}

	customer := model.Customer{}
	if err := c.Bind(&customer); err != nil {
		return util.WriteError(c, http.StatusBadRequest, "Bad request", err)
	}

	updatedCustomer, err := h.customerService.Update(uint(ID), &customer)
	if err != nil {
		log.Printf("Error updating customer: %v", err)
		return util.WriteError(c, http.StatusInternalServerError, "Failed to update customer", err)
	}

	return util.WriteResponse(c, http.StatusOK, "Customer updated successfully", updatedCustomer)
}

func (h *CustomerHandler) DeleteCustomer(c echo.Context) error {
	ID, err := util.ParseID(c)
	if err != nil {
		return util.WriteError(c, http.StatusBadRequest, "Invalid ID format", err)
	}

	if err := h.customerService.Delete(uint(ID)); err != nil {
		return util.WriteError(c, http.StatusInternalServerError, "Failed to delete customer", err)
	}

	return util.WriteResponse(c, http.StatusOK, "Customer deleted successfully", nil)
}
