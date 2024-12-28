package main

import (
	"fmt"
	"log"

	"github.com/IsraelTeo/api-store-go/authentication"
	"github.com/IsraelTeo/api-store-go/db"
	"github.com/IsraelTeo/api-store-go/handler"
	"github.com/IsraelTeo/api-store-go/repository"
	"github.com/IsraelTeo/api-store-go/route"
	"github.com/IsraelTeo/api-store-go/service"
	"github.com/IsraelTeo/api-store-go/validate"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	//Se crea repositorios, servicios y handlers.
	userRepository := repository.NewUserRepository(db.GDB)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	authentication := authentication.NewAuthService(userRepository)

	customerRepository := repository.NewCustomerRepository(db.GDB)
	customerService := service.NewCustomerService(customerRepository)
	customerHandler := handler.NewCustomerHandler(customerService)

	saleRepository := repository.NewSaleRepository(db.GDB)
	saleService := service.NewSaleRepository(saleRepository)
	saleHandler := handler.NewSaleHandler(saleService)

	producRepository := repository.NewProductRepository(db.GDB)
	producService := service.NewProductService(producRepository)
	productHandler := handler.NewProductHandler(producService)

	//Crea una nueva instancia del servidor web Echo.
	e := echo.New()

	//Instanciación de Rutas
	route.User(e, userHandler)
	route.Auth(e, authentication)
	route.Customer(e, customerHandler)
	route.Sale(e, saleHandler)
	route.Product(e, productHandler)

	//Inicialización del Validador
	validate.InitValidator()

	// Carga de variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loanding .env main")
	}

	//Este middleware agrega un logger personalizado para registrar el método HTTP, URI, estado y tiempo de latencia de cada solicitud.
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, time=${latency_human}\n",
	}))

	//Este middleware se asegura de que cualquier error de pánico en el servidor sea manejado correctamente y no detenga la ejecución.
	e.Use(middleware.Recover())

	//Registra cada solicitud HTTP de manera estándar.
	e.Use(middleware.Logger())

	err := db.Connection()
	if err != nil {
		log.Fatalf("Error trying to connect with database: %v", err)
	}
	fmt.Println("\nDatabase connection ok")

	err = db.MigrateDB()
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}
	fmt.Println("Database migration successful")

	fmt.Println("Starting server on port 8080...")

	err = e.Start(":8080")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}
