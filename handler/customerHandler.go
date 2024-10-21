package handler

import (
	"encoding/json"
	"net/http"

	"github.com/IsraelTeo/api-store-go/db"
	"github.com/IsraelTeo/api-store-go/model"
)

const contentTypeJSON = "application/json"

func setJSONContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", contentTypeJSON)
}

func create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		setJSONContentType(w)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type": "error", "message": "Method not permit"}`))
		return
	}

	customer := model.Customer{}
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		setJSONContentType(w)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type": "error", "message": "Invalid request body"}`))
		return
	}

	db.GDB.Create(&customer)
	setJSONContentType(w)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message_type": "success", "message": "Customer created successfully"}`))
}

func getCustomerById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		setJSONContentType(w)
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte(`{"message_type": "error", "message": "Method not permit"}`))
		return
	}
}
