package config

import (
	"log"

	"github.com/IsraelTeo/api-store-go/auth"
	authentication "github.com/IsraelTeo/api-store-go/auth"
	"github.com/IsraelTeo/api-store-go/handler"
	"github.com/labstack/echo/v4"
)

const (
	idPath   = "/:id"
	allPath  = "/all"
	voidPath = ""
)

func Run(e *echo.Echo,
	authService auth.LoginService,
	userHandler *handler.UserHandler,
	customerHandler *handler.CustomerHandler,
	saleHandler *handler.SaleHandler,
	productHandler *handler.ProductHandler) {

	e.Group("/api/v1") //Crea un enrutador

	// Grupo principal para api/v1
	api := e.Group("/api/v1")

	// Rutas de autenticación
	auth := api.Group("/auth")
	auth.POST("/login", authService.Login)

	// Rutas de usuarios
	users := api.Group("/users")
	users.POST(voidPath, userHandler.RegisterUser)
	users.GET(idPath, auth.ValidateJWT(userHandler.GetUserByID))
	users.GET(allPath, auth.ValidateJWT(userHandler.GetAllUsers))
	users.PUT(idPath, auth.ValidateJWT(userHandler.UpdateUser))
	users.DELETE(idPath, auth.ValidateJWT(userHandler.DeleteUser))

	// Rutas de clientes
	customers := api.Group("/customers")
	customers.POST(voidPath, auth.ValidateJWTAdmin(customerHandler.CreateCustomer))
	customers.GET(idPath, auth.ValidateJWTAdmin(customerHandler.GetCustomerByID))
	customers.GET(allPath, authentication.ValidateJWTAdmin(customerHandler.GetAllCustomers))
	customers.PUT(idPath, authentication.ValidateJWTAdmin(customerHandler.UpdateCustomer))
	customers.DELETE(idPath, authentication.ValidateJWTAdmin(customerHandler.DeleteCustomer))

	// Rutas de ventas
	sales := api.Group("/sales")
	sales.POST(voidPath, auth.ValidateJWTAdmin(saleHandler.CreateSale))
	sales.GET(idPath, auth.ValidateJWTAdmin(saleHandler.GetSaleByID))
	sales.GET(allPath, auth.ValidateJWTAdmin(saleHandler.GetAllSales))
	sales.PUT(idPath, auth.ValidateJWTAdmin(saleHandler.UpdateSale))
	sales.DELETE(idPath, auth.ValidateJWTAdmin(saleHandler.DeleteSale))

	// Rutas de productos
	products := api.Group("/products")
	products.POST(voidPath, auth.ValidateJWTAdmin(productHandler.CreateProduct))
	products.GET(idPath, auth.ValidateJWT(productHandler.GetProductByID))
	products.GET(allPath, auth.ValidateJWT(productHandler.GetAllProducts))
	products.PUT(idPath, auth.ValidateJWTAdmin(productHandler.UpdateProduct))
	products.DELETE(idPath, auth.ValidateJWTAdmin(productHandler.DeleteProduct))

	log.Println("Listening on:", s.addr)
}
