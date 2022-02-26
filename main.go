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

	AddRoutes(r)
	http.Handle("/", r)

	// Bind to a port and pass our router in
	db := database.GetDBConnection()
	fmt.Printf("Database Running on Port %d \n", database.Port)
	_ = db.Close()

	fmt.Println("Web server running on 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}
