package main

import (
	"backend/database"
	"backend/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"192.168.1.2"})

	// Test Connection with Database and Close it
	db := database.NewDBConnection()
	_ = db.Close()

	// Import Feature Vector from JSON
	database.AllFeatures.ImportFeaturesFromJSON()

	// Add Routes to our Routes
	router.GET("/", handlers.HomeHandler)
	v1 := router.Group("/v1")
	{
		v1.GET("/:lon/:lat/:radius", handlers.BoundingBoxHandler)
	}

	// Bind to a port and pass our router in
	fmt.Println("Web server running on 8080")
	log.Fatal(router.Run(":8080"))

}
