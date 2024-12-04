package handler

import (
	"encoding/json"
	"net/http"

	"github.com/IsraelTeo/api-store-go/db"
	"github.com/IsraelTeo/api-store-go/model"
	"github.com/gorilla/mux"
)

func GetSaleById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := newResponse(Error, "Method get not permitted", nil)
		responseJSON(w, http.StatusMethodNotAllowed, response)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	sale := model.Sale{}
	result := db.GDB.First(&sale, id)
	if result.Error != nil {
		response := newResponse(Error, "Sale are not found", nil)
		responseJSON(w, http.StatusNotFound, response)
		return
	}

	response := newResponse("success", "Sale found", sale)
	responseJSON(w, http.StatusOK, response)
}

func GetAllSales(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := newResponse(Error, "Method get not permitted", nil)
		responseJSON(w, http.StatusMethodNotAllowed, response)
	}

	var sales model.Sales
	result := db.GDB.Find(&sales)
	if result.Error != nil {
		response := newResponse(Error, "Failed to fetch customers", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	if len(sales) == 0 {
		response := newResponse("success", "Sales list is empty", sales)
		responseJSON(w, http.StatusNoContent, response)
		return
	}

	response := newResponse("success", "Sales found", sales)
	responseJSON(w, http.StatusOK, response)
}

func CreateSale(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := newResponse(Error, "Method post not permitted", nil)
		responseJSON(w, http.StatusMethodNotAllowed, response)
		return
	}

	var sale model.Sale
	result := db.GDB.Create(&sale)
	if result.Error != nil {
		response := newResponse(Error, "Bad request", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	response := newResponse("success", "Sale created successfusly", sale)
	responseJSON(w, http.StatusCreated, response)
}

func UpdateSale(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		response := newResponse(Error, "Method put not permitted", nil)
		responseJSON(w, http.StatusMethodNotAllowed, response)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	sale := model.Sale{}
	result := db.GDB.First(&sale, id)
	if result.Error != nil {
		response := newResponse(Error, "Sale not found", nil)
		responseJSON(w, http.StatusNotFound, response)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&sale)
	if err != nil {
		response := newResponse(Error, "Error decoding request body", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	db.GDB.Save(&sale)
	response := newResponse("success", "Sale updated successfully", sale)
	responseJSON(w, http.StatusOK, response)
}

func DeleteSale(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response := newResponse(Error, "Method delete not permit", nil)
		responseJSON(w, http.StatusMethodNotAllowed, response)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	sale := model.Sale{}
	result := db.GDB.First(&sale, id)
	if result.Error != nil {
		response := newResponse(Error, "Sale not found to delete", nil)
		responseJSON(w, http.StatusNotFound, response)
		return
	}

	db.GDB.Delete(&sale)
	response := newResponse("success", "Sale deleted successfull", nil)
	responseJSON(w, http.StatusOK, response)
}
