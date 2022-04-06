package handlers

import (
	"backend/app/controllers"
	"backend/app/database"
	"backend/app/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Handlers struct {
	Logger       *log.Logger
	DBConnection *database.DBConnection
}

func NewHandler(db *database.DBConnection, logger *log.Logger) *Handlers {
	return &Handlers{
		Logger:       logger,
		DBConnection: db,
	}
}

func (h *Handlers) Home(c *gin.Context) {
	h.Logger.Println("Accessing Home")
	c.Writer.Header().Set("Content-Type", "application/json")
	jsonObj := database.AllFeatures.Features
	c.JSON(http.StatusOK, jsonObj)
}

func (h *Handlers) BoundingBox(c *gin.Context) {
	h.Logger.Println("Creating BoundingBox for ", c.Param("lat"), c.Param("lon"), c.Param("radius"))
	c.Writer.Header().Set("Content-Type", "application/json")
	gpsObject, err := GeoDataController.ParseGeoRequest(c.Param("lat"), c.Param("lon"), c.Param("radius"))
	if err != nil {
		c.String(http.StatusBadRequest, "Error when trying to Parse values. Please Provide Numerical")
		return
	}

	payload := GeoDataController.GetFeatureVectors(gpsObject, database.NewDBConnection())
	c.JSON(http.StatusOK, payload)
}

func (h *Handlers) FrontendPass(c *gin.Context) {
	h.Logger.Println("Creating Frontend BoundingBox for ", c.Param("lat"), c.Param("lon"), c.Param("radius"))
	c.Writer.Header().Set("Content-Type", "application/json")
	gpsObject, err := GeoDataController.ParseGeoRequest(c.Param("lat"), c.Param("lon"), c.Param("radius"))
	if err != nil {
		c.String(http.StatusBadRequest, "Error when trying to Parse values. Please Provide Numerical")
		return
	}

	interim := GeoDataController.GetFeatureVectors(gpsObject, h.DBConnection)
	payload := models.PassthroughObj{FeatureVector: interim}
	c.JSON(http.StatusOK, payload)
}
