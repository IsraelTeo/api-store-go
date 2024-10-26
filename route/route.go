package route

import (
	"github.com/IsraelTeo/api-store-go/handler"
	"github.com/IsraelTeo/api-store-go/middelware"
	"github.com/gorilla/mux"
)

const (
	customersPath    = "/customers"
	customerIDPath   = "/customer/{id}"
	customerBasePath = "/customer"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	api := r.PathPrefix("/v1").Subrouter()

	api.HandleFunc(customersPath, middelware.Log(handler.GetAllCustomers)).Methods("GET")
	api.HandleFunc(customerIDPath, middelware.Log(handler.GetCustomerById)).Methods("GET")
	api.HandleFunc(customerBasePath, middelware.Log(middelware.Authentication(handler.CreateCustomer))).Methods("POST")
	api.HandleFunc(customerIDPath, middelware.Log(handler.UpdateCustomer)).Methods("PUT")
	api.HandleFunc(customerIDPath, middelware.Log(handler.DeleteCustomer)).Methods("DELETE")

	return r
}
