package main

import (
	"fmt"
	"log"

	"github.com/IsraelTeo/api-store-go/config"
	"github.com/IsraelTeo/api-store-go/db"
	"github.com/IsraelTeo/api-store-go/route"
	"github.com/IsraelTeo/api-store-go/util"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// Cargar las variables de entorno**
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Inicializar la configuración cargando las variables de entorno
	cfg := config.InitConfig()

	// Conectar a la base de datos utilizando la configuración cargada
	if err := db.Connection(cfg); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	fmt.Println("Database connection successful")

	//Migración de entidades
	if err := db.MigrateDB(); err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}
	fmt.Println("Database migration successful")

	// Inicializar servidor Echo
	e := echo.New()

	//Asignar el validador a la instancia de Echo
	e.Validator = util.InitValidator()

	//Instanciar Rutas
	route.RunRoutes(e)

	//Middlewares
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, time=${latency_human}\n",
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	//Inicia servidor en el puerto: 8080
	if err := config.StartServer(e, ":8080"); err != nil {
		log.Fatalf("%v", err)
	}

}
