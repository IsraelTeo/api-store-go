package route

import (
	"github.com/IsraelTeo/api-store-go/auth"
	"github.com/IsraelTeo/api-store-go/db"
	"github.com/IsraelTeo/api-store-go/handler"
	"github.com/IsraelTeo/api-store-go/repository"
	"github.com/IsraelTeo/api-store-go/service"
	"github.com/labstack/echo/v4"
)

// Constantes comunes para las rutas
const (
	idPath   = "/:id"
	allPath  = "/all"
	voidPath = ""
)

// RunRoutes configura las rutas principales de la aplicación
func RunRoutes(e *echo.Echo) {
	api := e.Group("/api/v1") // Crea un grupo principal de rutas

	// Rutas de autenticación
	setupAuthRoutes(api)

	// Rutas de usuarios
	setupUserRoutes(api)

	// Rutas de clientes
	setupCustomerRoutes(api)

	// Rutas de ventas
	setupSaleRoutes(api)

	// Rutas de productos
	setupProductRoutes(api)
}

// Configuración de rutas de autenticación
func setupAuthRoutes(api *echo.Group) {
	userRepository := repository.NewUserRepository(db.GDB)
	authService := auth.NewLogin(userRepository)

	authRoute := api.Group("/auth")
	authRoute.POST("/login", authService.Login)
}

// Configura las rutas de usuarios
func setupUserRoutes(api *echo.Group) {
	userRepository := repository.NewUserRepository(db.GDB)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	users := api.Group("/users")
	users.POST(voidPath, userHandler.RegisterUser)
	users.GET(idPath, auth.ValidateJWT(userHandler.GetUserByID))
	users.GET(allPath, auth.ValidateJWT(userHandler.GetAllUsers))
	users.PUT(idPath, auth.ValidateJWT(userHandler.UpdateUser))
	users.DELETE(idPath, auth.ValidateJWT(userHandler.DeleteUser))
}

// Configura las rutas de clientes
func setupCustomerRoutes(api *echo.Group) {
	customerRepository := repository.NewCustomerRepository(db.GDB)
	customerService := service.NewCustomerService(customerRepository)
	customerHandler := handler.NewCustomerHandler(customerService)

	customers := api.Group("/customers")
	customers.POST(voidPath, auth.ValidateJWTAdmin(customerHandler.CreateCustomer))
	customers.GET(idPath, auth.ValidateJWTAdmin(customerHandler.GetCustomerByID))
	customers.GET(allPath, auth.ValidateJWTAdmin(customerHandler.GetAllCustomers))
	customers.PUT(idPath, auth.ValidateJWTAdmin(customerHandler.UpdateCustomer))
	customers.DELETE(idPath, auth.ValidateJWTAdmin(customerHandler.DeleteCustomer))
}

// Configura las rutas de ventas
func setupSaleRoutes(api *echo.Group) {
	saleRepository := repository.NewSaleRepository(db.GDB)
	saleService := service.NewSaleRepository(saleRepository)
	saleHandler := handler.NewSaleHandler(saleService)

	sales := api.Group("/sales")
	sales.POST(voidPath, auth.ValidateJWTAdmin(saleHandler.CreateSale))
	sales.GET(idPath, auth.ValidateJWTAdmin(saleHandler.GetSaleByID))
	sales.GET(allPath, auth.ValidateJWTAdmin(saleHandler.GetAllSales))
	sales.PUT(idPath, auth.ValidateJWTAdmin(saleHandler.UpdateSale))
	sales.DELETE(idPath, auth.ValidateJWTAdmin(saleHandler.DeleteSale))
}

// Configura las rutas de productos
func setupProductRoutes(api *echo.Group) {
	productRepository := repository.NewProductRepository(db.GDB)
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	products := api.Group("/products")
	products.POST(voidPath, auth.ValidateJWTAdmin(productHandler.CreateProduct))
	products.GET(idPath, auth.ValidateJWT(productHandler.GetProductByID))
	products.GET(allPath, auth.ValidateJWT(productHandler.GetAllProducts))
	products.PUT(idPath, auth.ValidateJWTAdmin(productHandler.UpdateProduct))
	products.DELETE(idPath, auth.ValidateJWTAdmin(productHandler.DeleteProduct))
}
