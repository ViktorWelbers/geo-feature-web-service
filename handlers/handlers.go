package handlers

import (
	"backend/controllers"
	"backend/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func HomeHandler(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	jsonObj := database.AllFeatures.Features
	c.JSON(http.StatusOK, jsonObj)
}

func BoundingBoxHandler(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	lat, err1 := strconv.ParseFloat(c.Param("lat"), 64)
	lon, err2 := strconv.ParseFloat(c.Param("lon"), 64)
	radius, err3 := strconv.ParseFloat(c.Param("radius"), 64)
	if err1 != nil || err2 != nil || err3 != nil {
		c.String(http.StatusBadRequest, "Error when trying to Parse values. Please Provide Numerical")
		return
	}

	payload := GeoDataController.GetFeatureVectors(lat, lon, radius, database.NewDBConnection())

	c.JSON(http.StatusOK, payload)
}
