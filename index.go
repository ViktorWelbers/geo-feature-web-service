package main

import (
	"backend/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func routing(r *mux.Router) {
	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/products", handlers.ProductsHandler).Methods("GET")
	r.HandleFunc("/articles", handlers.ArticlesHandler).Methods("GET")
}

func main() {
	r := mux.NewRouter()
	routing(r)
	http.Handle("/", r)
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
