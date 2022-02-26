package handlers

import (
	"backend/controllers"
	"backend/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jsonObj := models.JsonResponse{Type: "Web-API", Data: "Empty", Message: "Server is running"}

	payload, err := json.Marshal(jsonObj)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Write(payload)
}

func BoundingBoxHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	lat, _ := strconv.ParseFloat(vars["lat"], 64)
	lon, _ := strconv.ParseFloat(vars["lon"], 64)
	radius, _ := strconv.ParseFloat(vars["radius"], 64)

	payload := GeoDataController.GetFeatureVectors(lat, lon, radius)

	w.Write(payload)
}

func ArticlesHandler(w http.ResponseWriter, r *http.Request) {
	println("cancer")
}
