package handler

import (
	"net/http"

	"github.com/IsraelTeo/api-store-go/model"
	"github.com/IsraelTeo/api-store-go/payload"
	"github.com/IsraelTeo/api-store-go/service"
	"github.com/IsraelTeo/api-store-go/validate"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

func (h *ProductHandler) GetProductByID(c echo.Context) error {
	ID, err := validate.ValidateAndParseID(c)
	if err != nil {
		response := payload.NewResponse(payload.MessageTypeError, "Invalid ID format.", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	product, err := h.productService.GetByID(uint(ID))
	if err != nil {
		response := payload.NewResponse(payload.MessageTypeError, "Product not found", nil)
		return c.JSON(http.StatusNotFound, response)
	}

	response := payload.NewResponse(payload.MessageTypeSuccess, "Product found", product)
	return c.JSON(http.StatusOK, response)
}

func (h *ProductHandler) GetAllProducts(c echo.Context) error {
	products, err := h.productService.GetAll()
	if err != nil {
		response := payload.NewResponse(payload.MessageTypeError, "Failed to fetch products", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := payload.NewResponse(payload.MessageTypeSuccess, "Products found", products)
	return c.JSON(http.StatusOK, response)
}

func (h *ProductHandler) CreateProduct(c echo.Context) error {
	product := model.Product{}
	if err := c.Bind(&product); err != nil {
		response := payload.NewResponse(payload.MessageTypeError, "Bad request", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := h.productService.Create(&product); err != nil {
		response := payload.NewResponse(payload.MessageTypeError, "Failed to save product", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := payload.NewResponse(payload.MessageTypeSuccess, "Product created successfully", nil)
	return c.JSON(http.StatusCreated, response)
}

func (h *ProductHandler) UpdateProduct(c echo.Context) error {
	ID, err := validate.ValidateAndParseID(c)
	if err != nil {
		response := payload.NewResponse(payload.MessageTypeError, "Invalid ID format", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	product := model.Product{}
	if err := c.Bind(&product); err != nil {
		response := payload.NewResponse(payload.MessageTypeError, "Bad request", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	updatedProduct, err := h.productService.Update(uint(ID), &product)
	if err != nil {
		response := payload.NewResponse(payload.MessageTypeError, "Failed to update product", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := payload.NewResponse(payload.MessageTypeSuccess, "Product updated successfully", updatedProduct)
	return c.JSON(http.StatusOK, response)
}

func (h *ProductHandler) DeleteProduct(c echo.Context) error {
	ID, err := validate.ValidateAndParseID(c)
	if err != nil {
		response := payload.NewResponse(payload.MessageTypeError, "Invalid ID format", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := h.productService.Delete(uint(ID)); err != nil {
		response := payload.NewResponse(payload.MessageTypeError, "Failed to delete product", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := payload.NewResponse(payload.MessageTypeSuccess, "Product deleted successfully", nil)
	return c.JSON(http.StatusOK, response)
}
