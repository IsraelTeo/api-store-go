package main

import (
	"fmt"
	"log"
	"os"

	"github.com/IsraelTeo/api-store-go/config"
	"github.com/IsraelTeo/api-store-go/db"
	"github.com/IsraelTeo/api-store-go/route"
	"github.com/IsraelTeo/api-store-go/validate"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// 1️⃣ **Cargar variables de entorno**
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// 2️⃣ **Conexión a la base de datos**
	if err := db.Connection(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	fmt.Println("✅ Database connection successful")

	// 3️⃣ **Migración de entidades**
	if err := db.MigrateDB(); err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}
	fmt.Println("✅ Database migration successful")

	// 4️⃣ **Inicializar servidor Echo**
	e := echo.New()

	// 🔑 Asignar el validador a la instancia de Echo
	e.Validator = validate.InitValidator()

	// Instanciar Rutas
	route.RunRoutes(e)

	// 6️⃣ **Middlewares**
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, time=${latency_human}\n",
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// 8️⃣ **Iniciar Servidor**
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	config.StartServer(e, ":"+port)
	fmt.Printf("🚀 Starting server on port %s...\n", port)
}
