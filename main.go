package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/IsraelTeo/api-store-go/authorization"
	"github.com/IsraelTeo/api-store-go/db"
	"github.com/IsraelTeo/api-store-go/route"
	"github.com/gorilla/mux"
)

func main() {

	err := authorization.LoadFiles("certificates/appp.rsa", "certificates.app")
	if err != nil {
		log.Fatalf("cant not load the certificates: %v", err)
	}

	r := mux.NewRouter()

	err = db.Connection()
	if err != nil {
		log.Fatalf("Error trying to connect with database: %v", err)
	}
	fmt.Println("Database connection ok")

	err = db.MigrateDB()
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}
	fmt.Println("Database migration successful")

	fmt.Println("Starting server on port 8080...")

	r = route.SetupRoutes()

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}
