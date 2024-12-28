package route

import (
	"github.com/IsraelTeo/api-store-go/authentication"
	"github.com/IsraelTeo/api-store-go/handler"
	"github.com/labstack/echo/v4"
)

const (
	idPath   = "/:id"
	allPath  = "/all"
	voidPath = ""
)

func SetupRoutes(e *echo.Echo,
	authService authentication.LoginService,
	userHandler *handler.UserHandler,
	customerHandler *handler.CustomerHandler,
	saleHandler *handler.SaleHandler,
	productHandler *handler.ProductHandler) {

	// Grupo principal para api/v1
	api := e.Group("/api/v1")

	// Rutas de autenticación
	auth := api.Group("/auth")
	auth.POST("/login", authService.Login)

	// Rutas de usuarios
	users := api.Group("/users")
	users.POST(voidPath, userHandler.RegisterUser)
	users.GET(idPath, authentication.ValidateJWT(userHandler.GetUserByID))
	users.GET(allPath, authentication.ValidateJWT(userHandler.GetAllUsers))
	users.PUT(idPath, authentication.ValidateJWT(userHandler.UpdateUser))
	users.DELETE(idPath, authentication.ValidateJWT(userHandler.DeleteUser))

	// Rutas de clientes
	customers := api.Group("/customers")
	customers.POST(voidPath, authentication.ValidateJWTAdmin(customerHandler.CreateCustomer))
	customers.GET(idPath, authentication.ValidateJWTAdmin(customerHandler.GetCustomerByID))
	customers.GET(allPath, authentication.ValidateJWTAdmin(customerHandler.GetAllCustomers))
	customers.PUT(idPath, authentication.ValidateJWTAdmin(customerHandler.UpdateCustomer))
	customers.DELETE(idPath, authentication.ValidateJWTAdmin(customerHandler.DeleteCustomer))

	// Rutas de ventas
	sales := api.Group("/sales")
	sales.POST(voidPath, saleHandler.CreateSale)
	sales.GET(idPath, saleHandler.GetSaleByID)
	sales.GET(allPath, saleHandler.GetAllSales)
	sales.PUT(idPath, saleHandler.UpdateSale)
	sales.DELETE(idPath, saleHandler.DeleteSale)

	// Rutas de productos
	products := api.Group("/products")
	products.POST(voidPath, productHandler.CreateProduct)
	products.GET(idPath, authentication.ValidateJWT(productHandler.GetProductByID))
	products.GET(allPath, authentication.ValidateJWT(productHandler.GetAllProducts))
	products.PUT(idPath, productHandler.UpdateProduct)
	products.DELETE(idPath, productHandler.DeleteProduct)
}
