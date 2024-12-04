package handler

import (
	"encoding/json"
	"net/http"

	"github.com/IsraelTeo/api-store-go/db"
	"github.com/IsraelTeo/api-store-go/model"
	"github.com/gorilla/mux"
)

func GetProductById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := newResponse(Error, "Method get not permitted", nil)
		responseJSON(w, http.StatusMethodNotAllowed, response)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	sale := model.Sale{}
	result := db.GDB.First(id, &sale)
	if result.Error != nil {
		response := newResponse(Error, "Sale not found", nil)
		responseJSON(w, http.StatusNotFound, response)
		return
	}

	response := newResponse("success", "Product found", sale)
	responseJSON(w, http.StatusOK, response)
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := newResponse(Error, "Method get not permitted", nil)
		responseJSON(w, http.StatusMethodNotAllowed, response)
		return
	}

	var products model.Products
	result := db.GDB.Find(&products)
	if result.Error != nil {
		response := newResponse(Error, "Failed to fetch products", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	if len(products) == 0 {
		response := newResponse("success", "Products list is empty", nil)
		responseJSON(w, http.StatusNoContent, response)
		return
	}

	response := newResponse("success", "Products found", products)
	responseJSON(w, http.StatusOK, response)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := newResponse(Error, "Method post not permitted", nil)
		responseJSON(w, http.StatusMethodNotAllowed, response)
		return
	}

	var product model.Product
	result := db.GDB.Create(&product)
	if result.Error != nil {
		response := newResponse(Error, "Product not found", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	response := newResponse("success", "Product created successfusly", product)
	responseJSON(w, http.StatusOK, response)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		response := newResponse(Error, "Method put not permitted", nil)
		responseJSON(w, http.StatusMethodNotAllowed, response)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	product := model.Product{}

	result := db.GDB.First(&product, id)
	if result.Error != nil {
		response := newResponse(Error, "Product not found", nil)
		responseJSON(w, http.StatusNotFound, response)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		response := newResponse(Error, "Error decoding request body", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	db.GDB.Save(&product)
	response := newResponse("success", "Product updated successfully", product)
	responseJSON(w, http.StatusOK, response)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response := newResponse(Error, "Method delete not permit", nil)
		responseJSON(w, http.StatusMethodNotAllowed, response)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	product := model.Product{}
	result := db.GDB.First(&product, id)
	if result.Error != nil {
		response := newResponse(Error, "Product not found to delete", nil)
		responseJSON(w, http.StatusNotFound, response)
		return
	}

	db.GDB.Delete(&product)
	response := newResponse("success", "Product deleted successfull", nil)
	responseJSON(w, http.StatusOK, response)
}
