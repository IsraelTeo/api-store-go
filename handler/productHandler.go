package handler

import (
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

	response := newResponse("success", "Sale found", sale)
	responseJSON(w, http.StatusOK, response)
}


func GetAllProducts(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet{
		response := newResponse(Error, "Method get not permitted", nil)
		responseJSON(w, http.StatusMethodNotAllowed, response)
		return
	}

	var sales []model.Sale
	result := db.GDB.Find(&sales)
	if result.Error != nil{
		response := newResponse(Error, "Sales not found", nil)
		responseJSON(w, http.StatusNotFound, response)
		return
	}

	response := newResponse("success", "Sales found", sales)
	responseJSON(w, http.StatusOK, response)
}









}
