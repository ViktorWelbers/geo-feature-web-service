package main

import (
	"backend/handlers"
	"backend/repositories"
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

	// Import Feature Vector from JSON-File
	repositories.AllFeatures.ImportFeaturesFromJSON()

	// Add Routes to our Routes
	router.GET("/", handlers.NewHandler(nil, logger).Home)
	v1 := router.Group("/v1")
	{
		v1.GET("/:lon/:lat/:radius", handlers.NewHandler(repositories.NewDBConnection(), logger).BoundingBox)
	}
	v2 := router.Group("/v2")
	{
		v2.GET("/:lon/:lat/:radius", handlers.NewHandler(repositories.NewDBConnection(), logger).FrontendPass)
	}
	// Bind to a port and pass our router in
	logger.Println("Web server running on 8080")
	logger.Fatal(router.Run(":8080"))

}
