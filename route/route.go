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

func Auth(e *echo.Echo, authentication authentication.LoginService) {
	auth := e.Group("/auth")
	auth.POST("/login", authentication.Login)
}

func User(e *echo.Echo, userHandler *handler.UserHandler) {
	users := e.Group("/users")
	users.POST("", userHandler.RegisterUser)
	users.GET(idPath, userHandler.GetUserByID)
	users.GET(allPath, userHandler.GetAllUsers)
	users.PUT(idPath, userHandler.UpdateUser)
	users.DELETE(idPath, userHandler.DeleteUser)
}

func Customer(e *echo.Echo, customerHandler *handler.CustomerHandler) {
	customers := e.Group("/customers")
	customers.POST("", customerHandler.CreateCustomer)
	customers.GET(idPath, customerHandler.GetCustomerByID)
	customers.GET(allPath, customerHandler.GetAllCustomers)
	customers.PUT(idPath, customerHandler.UpdateCustomer)
	customers.DELETE(idPath, customerHandler.DeleteCustomer)
}

func Sale(e *echo.Echo, saleHandler *handler.SaleHandler) {
	sales := e.Group("/sales")
	sales.POST("", saleHandler.CreateSale)
	sales.GET(idPath, saleHandler.GetSaleByID)
	sales.GET(allPath, saleHandler.GetAllSales)
	sales.PUT(idPath, saleHandler.UpdateSale)
	sales.DELETE(idPath, saleHandler.DeleteSale)
}

func Product(e *echo.Echo, productHandler *handler.ProductHandler) {
	products := e.Group("/products")
	products.POST("", productHandler.CreateProduct)
	products.GET(idPath, productHandler.GetProductByID)
	products.GET(allPath, productHandler.GetAllProducts)
	products.PUT(idPath, productHandler.UpdateProduct)
	products.DELETE(idPath, productHandler.DeleteProduct)
}
