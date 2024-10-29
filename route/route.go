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

	productsPath    = "/products"
	productIDPath   = "/product/{id}"
	productBasePath = "/product"

	salesPath    = "/sales"
	saleIDPath   = "/sale/{id}"
	saleBasePath = "/sale"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	api := r.PathPrefix("/v1").Subrouter()

	api.HandleFunc(customersPath, middelware.Log(handler.GetAllCustomers)).Methods("GET")
	api.HandleFunc(customerIDPath, middelware.Log(handler.GetCustomerById)).Methods("GET")
	api.HandleFunc(customerBasePath, middelware.Log(handler.CreateCustomer)).Methods("POST")
	api.HandleFunc(customerIDPath, middelware.Log(handler.UpdateCustomer)).Methods("PUT")
	api.HandleFunc(customerIDPath, middelware.Log(handler.DeleteCustomer)).Methods("DELETE")

	api.HandleFunc(productsPath, middelware.Log(handler.GetAllProducts)).Methods("GET")
	api.HandleFunc(productIDPath, middelware.Log(handler.GetProductById)).Methods("GET")
	api.HandleFunc(productBasePath, middelware.Log(handler.CreateProduct)).Methods("POST")
	api.HandleFunc(productIDPath, middelware.Log(handler.UpdateProduct)).Methods("PUT")
	api.HandleFunc(productIDPath, middelware.Log(handler.DeleteProduct)).Methods("DELETE")

	api.HandleFunc(salesPath, middelware.Log(handler.GetAllSales)).Methods("GET")
	api.HandleFunc(saleIDPath, middelware.Log(handler.GetSaleById)).Methods("GET")
	api.HandleFunc(saleBasePath, middelware.Log(handler.CreateSale)).Methods("POST")
	api.HandleFunc(saleIDPath, middelware.Log(handler.UpdateSale)).Methods("PUT")
	api.HandleFunc(saleIDPath, middelware.Log(handler.DeleteSale)).Methods("DELETE")

	return r
}
