package handler

import (
	"net/http"

	"github.com/IsraelTeo/api-store-go/model"
	"github.com/IsraelTeo/api-store-go/service"
	"github.com/IsraelTeo/api-store-go/util"
	"github.com/labstack/echo/v4"
)

type SaleHandler struct {
	service service.SaleService
}

func NewSaleHandler(service service.SaleService) *SaleHandler {
	return &SaleHandler{service: service}
}

func (h *SaleHandler) GetSaleByID(c echo.Context) error {
	ID, err := util.ParseID(c)
	if err != nil {
		return util.WriteError(c, http.StatusBadRequest, "Invalid ID", nil)
	}

	sale, err := h.service.GetByID(uint(ID))
	if err != nil {
		return util.WriteError(c, http.StatusNotFound, "Sale not found", nil)
	}

	return util.WriteResponse(c, http.StatusOK, "Sale found successfully", sale)
}

func (h *SaleHandler) GetAllSales(c echo.Context) error {
	sales, err := h.service.GetAll()
	if err != nil {
		return util.WriteError(c, http.StatusInternalServerError, "Failed to fetch sales", err)
	}

	return util.WriteResponse(c, http.StatusOK, "Sales found", sales)
}

func (h *SaleHandler) CreateSale(c echo.Context) error {
	sale := model.Sale{}
	if err := c.Bind(&sale); err != nil {
		return util.WriteError(c, http.StatusBadRequest, "Bad request", err)
	}

	if err := h.service.Create(&sale); err != nil {
		return util.WriteError(c, http.StatusInternalServerError, "Failed to save sale", err)
	}

	return util.WriteResponse(c, http.StatusCreated, "Sale created successfully", sale)
}

// UpdateSale actualiza una venta existente
func (h *SaleHandler) UpdateSale(c echo.Context) error {
	ID, err := util.ParseID(c)
	if err != nil {
		return util.WriteError(c, http.StatusBadRequest, "Invalid ID format", err)
	}

	sale := model.Sale{}
	if err := c.Bind(&sale); err != nil {
		return util.WriteError(c, http.StatusBadRequest, "Bad request", err)
	}

	updatedSale, err := h.service.Update(uint(ID), &sale)
	if err != nil {
		return util.WriteError(c, http.StatusInternalServerError, "Failed to update sale", err)
	}

	return util.WriteResponse(c, http.StatusOK, "Sale updated successfully", updatedSale)
}

func (h *SaleHandler) DeleteSale(c echo.Context) error {
	ID, err := util.ParseID(c)
	if err != nil {
		return util.WriteError(c, http.StatusBadRequest, "Invalid ID format", err)
	}

	if err := h.service.Delete(uint(ID)); err != nil {
		return util.WriteError(c, http.StatusInternalServerError, "Failed to delete sale", err)
	}

	return util.WriteResponse(c, http.StatusOK, "Sale deleted successfully", nil)
}
