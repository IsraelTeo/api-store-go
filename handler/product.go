package handler

import (
	"net/http"

	"github.com/IsraelTeo/api-store-go/model"
	"github.com/IsraelTeo/api-store-go/service"
	"github.com/IsraelTeo/api-store-go/util"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

func (h *ProductHandler) GetProductByID(c echo.Context) error {
	ID, err := util.ParseID(c)
	if err != nil {
		return util.WriteError(c, http.StatusBadRequest, "Invalid ID", nil)
	}

	product, err := h.productService.GetByID(uint(ID))
	if err != nil {
		return util.WriteError(c, http.StatusNotFound, "Product not found", nil)
	}

	return util.WriteResponse(c, http.StatusOK, "Product found", product)
}

func (h *ProductHandler) GetAllProducts(c echo.Context) error {
	products, err := h.productService.GetAll()
	if err != nil {
		return util.WriteError(c, http.StatusInternalServerError, "Failed to fetch products", err)
	}

	return util.WriteResponse(c, http.StatusOK, "Products found", products)
}

func (h *ProductHandler) CreateProduct(c echo.Context) error {
	product := model.Product{}
	if err := c.Bind(&product); err != nil {
		return util.WriteError(c, http.StatusBadRequest, "Bad request", err)
	}

	if err := h.productService.Create(&product); err != nil {
		return util.WriteError(c, http.StatusInternalServerError, "Failed to save product", err)
	}

	return util.WriteResponse(c, http.StatusCreated, "Product created successfully", nil)
}

func (h *ProductHandler) UpdateProduct(c echo.Context) error {
	ID, err := util.ParseID(c)
	if err != nil {
		return util.WriteError(c, http.StatusBadRequest, "Invalid ID format", err)
	}

	product := model.Product{}
	if err := c.Bind(&product); err != nil {
		return util.WriteError(c, http.StatusBadRequest, "Bad request", err)
	}

	updatedProduct, err := h.productService.Update(uint(ID), &product)
	if err != nil {
		return util.WriteError(c, http.StatusInternalServerError, "Failed to update product", err)
	}

	return util.WriteResponse(c, http.StatusOK, "Product updated successfully", updatedProduct)
}

func (h *ProductHandler) DeleteProduct(c echo.Context) error {
	ID, err := util.ParseID(c)
	if err != nil {
		return util.WriteError(c, http.StatusBadRequest, "Invalid ID format", err)
	}

	if err := h.productService.Delete(uint(ID)); err != nil {
		return util.WriteError(c, http.StatusInternalServerError, "Failed to delete product", err)
	}

	return util.WriteResponse(c, http.StatusOK, "Product deleted successfully", nil)
}
