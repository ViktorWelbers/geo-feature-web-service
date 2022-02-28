package main

import (
	"backend/database"
	"backend/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	// Test Connection with Database and Close it
	db, err := database.GetDBConnection()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	_ = db.Close()
	fmt.Printf("Database Running on Port %d \n", database.Port)

	// Import Feature Vector from JSON
	database.AllFeatures.ImportFeaturesFromJSON()

	// Add Routes to our Routes
	router.HandleFunc("/", handlers.HomeHandler)
	router.HandleFunc("/get-features/{lon}/{lat}/{radius}", handlers.BoundingBoxHandler).Methods("GET")

	// Bind to a port and pass our router in
	fmt.Println("Web server running on 8000")
	log.Fatal(http.ListenAndServe(":8000", router))

}
