package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/igor570/eventexplorer/db"
	"github.com/igor570/eventexplorer/routes"
)

func main() {
	// Initialize the database
	db.InitDB()

	// Create a Gin router with default middleware
	app := gin.Default()

	// Set up CORS middleware configuration
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	// Apply the CORS middleware to the router
	app.Use(cors.New(config))

	// Register routes
	routes.RegisterRoutes(app)

	// Start the server
	err := app.Run(":3100") // localhost:3100
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
