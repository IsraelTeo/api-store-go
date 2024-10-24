package route

import (
	"github.com/IsraelTeo/api-store-go/handler"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/customers", handler.GetAllCustomers).Methods("GET")
	r.HandleFunc("/customer/{id}", handler.GetCustomerById).Methods("GET")
	r.HandleFunc("/customer", handler.CreateCustomer).Methods("POST")
	r.HandleFunc("/customer/{id}", handler.UpdateCustomer).Methods("PUT")
	r.HandleFunc("/customer/{id}", handler.DeleteCustomer).Methods("DELETE")

	return r
}
