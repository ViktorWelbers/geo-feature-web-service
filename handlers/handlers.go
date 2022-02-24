package handlers

import (
	"backend/database"
	"encoding/json"
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jsonObj := database.JsonResponse{Type: "aids", Data: "smelly", Message: "Du kleiner Hurensohn"}
	payload, err := json.Marshal(jsonObj)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(payload)
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	println("cancer")
}

func ArticlesHandler(w http.ResponseWriter, r *http.Request) {
	println("cancer")
}
