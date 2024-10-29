package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/IsraelTeo/api-store-go/db"
	"github.com/IsraelTeo/api-store-go/model"
	"github.com/gorilla/mux"
)

func GetCustomerById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := newResponse(Error, "Method get not permitted", nil)
		responseJSON(w, http.StatusMethodNotAllowed, response)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	customer := model.Customer{}
	result := db.GDB.First(&customer, id)
	if result.Error != nil {
		response := newResponse(Error, "Customer not found", nil)
		responseJSON(w, http.StatusNotFound, response)
		return
	}

	response := newResponse("success", "Customer found", customer)
	responseJSON(w, http.StatusOK, response)
}

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := newResponse(Error, "Method GET not permitted", nil)
		responseJSON(w, http.StatusMethodNotAllowed, response)
		return
	}

	var customers []model.Customer
	result := db.GDB.Find(&customers)
	if result.Error != nil {
		response := newResponse(Error, "Failed to fetch customers", nil)
		fmt.Println("Error fetching customers:", result.Error)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	if len(customers) == 0 {
		response := newResponse("success", "Customer list is empty", nil)
		responseJSON(w, http.StatusNoContent, response)
		return
	}

	response := newResponse("success", "Customers found", customers)
	responseJSON(w, http.StatusOK, response)
}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := newResponse(Error, "Method post not permit", nil)
		responseJSON(w, http.StatusMethodNotAllowed, response)
		return
	}

	customer := model.Customer{}

	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		response := newResponse(Error, "Internal server error", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	db.GDB.Create(&customer)
	response := newResponse("success", "Customer created successfully", nil)
	responseJSON(w, http.StatusCreated, response)
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		response := newResponse(Error, "Method put not permitted", nil)
		responseJSON(w, http.StatusMethodNotAllowed, response)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	customer := model.Customer{}

	result := db.GDB.First(&customer, id)
	if result.Error != nil {
		response := newResponse(Error, "Customer not found", nil)
		responseJSON(w, http.StatusNotFound, response)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		response := newResponse(Error, "Error decoding request body", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	db.GDB.Save(&customer)
	response := newResponse("success", "Customer updated successfully", customer)
	responseJSON(w, http.StatusOK, response)
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response := newResponse(Error, "Method delete not permit", nil)
		responseJSON(w, http.StatusMethodNotAllowed, response)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	customer := model.Customer{}
	result := db.GDB.First(&customer, id)
	if result.Error != nil {
		response := newResponse(Error, "Customer not found to delete", nil)
		responseJSON(w, http.StatusNotFound, response)
		return
	}

	db.GDB.Delete(&customer)
	response := newResponse("success", "Customer deleted successfull", nil)
	responseJSON(w, http.StatusOK, response)
}
