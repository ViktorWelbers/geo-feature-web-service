package main

import (
	"backend/database"
	"backend/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func AddRoutes(r *mux.Router) {
	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/get-features/{lon}/{lat}/{radius}", handlers.BoundingBoxHandler).Methods("GET")
}

func main() {
	r := mux.NewRouter()

	// Test Connection with Database and Close it
	db := database.GetDBConnection()
	_ = db.Close()
	fmt.Printf("Database Running on Port %d \n", database.Port)

	// Import Feature Vector from JSON
	database.AllFeatures.ImportFeaturesFromJSON()

	// Add Routes to our Router
	http.Handle("/", r)
	AddRoutes(r)

	// Bind to a port and pass our router in
	fmt.Println("Web server running on 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}
