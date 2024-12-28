package route

import (
	"github.com/IsraelTeo/api-store-go/authentication"
	"github.com/IsraelTeo/api-store-go/handler"
	"github.com/labstack/echo/v4"
)

const (
	idPath  = "/:id"
	allPath = "/all"
)

func SetupRoutes(e *echo.Echo,
	authentication authentication.LoginService,
	userHandler *handler.UserHandler,
	customerHandler *handler.CustomerHandler,
	saleHandler *handler.SaleHandler,
	productHandler *handler.ProductHandler) {

	api := e.Group("/api/v1") // Grupo principal para api/v1

	// Rutas de autenticación
	auth := api.Group("/auth")
	auth.POST("/login", authentication.Login)

	// Rutas de usuarios
	users := api.Group("/users")
	users.POST("", userHandler.RegisterUser)
	users.GET(idPath, userHandler.GetUserByID)
	users.GET(allPath, userHandler.GetAllUsers)
	users.PUT(idPath, userHandler.UpdateUser)
	users.DELETE(idPath, userHandler.DeleteUser)

	// Rutas de clientes
	customers := api.Group("/customers")
	customers.POST("", customerHandler.CreateCustomer)
	customers.GET(idPath, customerHandler.GetCustomerByID)
	customers.GET(allPath, customerHandler.GetAllCustomers)
	customers.PUT(idPath, customerHandler.UpdateCustomer)
	customers.DELETE(idPath, customerHandler.DeleteCustomer)

	// Rutas de ventas
	sales := api.Group("/sales")
	sales.POST("", saleHandler.CreateSale)
	sales.GET(idPath, saleHandler.GetSaleByID)
	sales.GET(allPath, saleHandler.GetAllSales)
	sales.PUT(idPath, saleHandler.UpdateSale)
	sales.DELETE(idPath, saleHandler.DeleteSale)

	// Rutas de productos
	products := api.Group("/products")
	products.POST("", productHandler.CreateProduct)
	products.GET(idPath, productHandler.GetProductByID)
	products.GET(allPath, productHandler.GetAllProducts)
	products.PUT(idPath, productHandler.UpdateProduct)
	products.DELETE(idPath, productHandler.DeleteProduct)
}
