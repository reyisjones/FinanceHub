package main

import (
	"log"
	"os"

	"financehub/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Initialize Gin router
	router := gin.Default()

	// Configure CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173", "http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	router.Use(cors.New(config))

	// Initialize handlers
	h := handlers.NewHandler()

	// API routes
	api := router.Group("/api")
	{
		// Health check
		api.GET("/health", h.HealthCheck)

		// Topics
		api.GET("/topics", h.GetAllTopics)
		api.GET("/topics/:id", h.GetTopicByID)

		// Stocks
		api.GET("/stocks/:symbol", h.GetStockQuote)
		api.GET("/stocks/:symbol/timeseries", h.GetStockTimeSeries)

		// Cryptocurrencies
		api.GET("/crypto/top", h.GetTopCryptos)
		api.GET("/crypto/:id", h.GetCryptoPrice)

		// Currency Exchange
		api.GET("/currency/:from/:to", h.GetCurrencyRate)
	}

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	log.Printf("Finance Hub API server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
