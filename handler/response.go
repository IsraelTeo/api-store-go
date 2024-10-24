package handler

import (
	"encoding/json"
	"net/http"
)

const (
	Error   = "error"
	Message = "message"
)

type response struct {
	MessageType string      `json:"message_type"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
}

func newResponse(messageType, message string, data interface{}) response {
	return response{
		messageType,
		message,
		data,
	}
}

func responseJSON(w http.ResponseWriter, statusCode int, rep response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(&rep)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
}
