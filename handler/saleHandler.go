package handler

import (
	"net/http"

	"github.com/IsraelTeo/api-store-go/model"
	"github.com/IsraelTeo/api-store-go/payload"
	"github.com/IsraelTeo/api-store-go/service"
	"github.com/IsraelTeo/api-store-go/validate"
	"github.com/labstack/echo/v4"
)

type SaleHandler struct {
	service service.SaleService
}

// NewSaleHandler constructor para SaleHandler
func NewSaleHandler(service service.SaleService) *SaleHandler {
	return &SaleHandler{service: service}
}

// GetSaleById obtiene una venta por su ID
func (h *SaleHandler) GetSaleByID(c echo.Context) error {
	ID, err := validate.ValidateAndParseID(c)
	if err != nil {
		response := payload.NewResponse(payload.MessageTypeError, "Invalid ID format.", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	sale, err := h.service.GetByID(uint(ID))
	if err != nil {
		response := payload.NewResponse(payload.MessageTypeError, "Sale not found", nil)
		return c.JSON(http.StatusNotFound, response)
	}

	response := payload.NewResponse(payload.MessageTypeSuccess, "Sale found successfully", sale)
	return c.JSON(http.StatusOK, response)
}

// GetAllSales obtiene todas las ventas
func (h *SaleHandler) GetAllSales(c echo.Context) error {
	sales, err := h.service.GetAll()
	if err != nil {
		response := payload.NewResponse(payload.MessageTypeError, "Failed to fetch sales", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	if len(sales) == 0 {
		response := payload.NewResponse(payload.MessageTypeSuccess, "Sales list is empty", sales)
		return c.JSON(http.StatusNoContent, response)
	}

	response := payload.NewResponse(payload.MessageTypeSuccess, "Sales found", sales)
	return c.JSON(http.StatusOK, response)
}

// CreateSale registra una nueva venta
func (h *SaleHandler) CreateSale(c echo.Context) error {
	sale := model.Sale{}
	if err := c.Bind(&sale); err != nil {
		response := payload.NewResponse(payload.MessageTypeError, "Bad request", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := h.service.Create(&sale); err != nil {
		response := payload.NewResponse(payload.MessageTypeError, "Failed to save sale", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := payload.NewResponse(payload.MessageTypeSuccess, "Sale created successfully", sale)
	return c.JSON(http.StatusCreated, response)
}

// UpdateSale actualiza una venta existente
func (h *SaleHandler) UpdateSale(c echo.Context) error {
	ID, err := validate.ValidateAndParseID(c)
	if err != nil {
		response := payload.NewResponse(payload.MessageTypeError, "Invalid ID format", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	sale := model.Sale{}
	if err := c.Bind(&sale); err != nil {
		response := payload.NewResponse(payload.MessageTypeError, "Bad request", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	updatedSale, err := h.service.Update(uint(ID), &sale)
	if err != nil {
		response := payload.NewResponse(payload.MessageTypeError, "Failed to update sale", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := payload.NewResponse(payload.MessageTypeSuccess, "Sale updated successfully", updatedSale)
	return c.JSON(http.StatusOK, response)
}

// DeleteSale elimina una venta por su ID
func (h *SaleHandler) DeleteSale(c echo.Context) error {
	ID, err := validate.ValidateAndParseID(c)
	if err != nil {
		response := payload.NewResponse(payload.MessageTypeError, "Invalid ID format", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := h.service.Delete(uint(ID)); err != nil {
		response := payload.NewResponse(payload.MessageTypeError, "Failed to delete sale", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := payload.NewResponse(payload.MessageTypeSuccess, "Sale deleted successfully", nil)
	return c.JSON(http.StatusOK, response)
}
