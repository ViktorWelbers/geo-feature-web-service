package handlers

import (
	"backend/controllers"
	"backend/database"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jsonObj := database.AllFeatures.Features
	payload, err := json.Marshal(jsonObj)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Write(payload)
}

func BoundingBoxHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ACCEPT GET /get-features/{lon}/{lat}/{radius}", time.Now().Format("01-02-2006 15:04:05"))
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	lat, err1 := strconv.ParseFloat(vars["lat"], 64)
	lon, err2 := strconv.ParseFloat(vars["lon"], 64)
	radius, err3 := strconv.ParseFloat(vars["radius"], 64)
	if err1 != nil || err2 != nil || err3 != nil {
		http.Error(w, "One of the given values is not numeric", http.StatusBadRequest)
		return
	}
	payload := GeoDataController.GetFeatureVectors(lat, lon, radius)

	w.Write(payload)
	fmt.Println("COMPLETE GET /get-features/{lon}/{lat}/{radius}", time.Now().Format("01-02-2006 15:04:05"))
}
