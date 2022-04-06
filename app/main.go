package main

import (
	"backend/app/database"
	"backend/app/handlers"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	router := gin.Default()
	logger := log.New(os.Stdout, "gin-logger ", log.LstdFlags|log.Lshortfile)
	err := router.SetTrustedProxies([]string{"192.168.1.2"})
	if err != nil {
		logger.Fatal(err)
	}

	// Test Connection with Database and Close it
	db := database.NewDBConnection()
	_ = db.Close()
	// Import Feature Vector from JSON
	database.AllFeatures.ImportFeaturesFromJSON()

	// Add Routes to our Routes
	router.GET("/", handlers.NewHandler(nil, logger).Home)
	v1 := router.Group("/v1")
	{
		v1.GET("/:lon/:lat/:radius", handlers.NewHandler(database.NewDBConnection(), logger).BoundingBox)
	}
	v2 := router.Group("/v2")
	{
		v2.GET("/:lon/:lat/:radius", handlers.NewHandler(database.NewDBConnection(), logger).FrontendPass)
	}
	// Bind to a port and pass our router in
	logger.Println("Web server running on 8080")
	logger.Fatal(router.Run(":8080"))

}
