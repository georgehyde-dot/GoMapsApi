package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/georgehyde-dot/GoMapsApi/pkg/mapsapi"
	"github.com/joho/godotenv"
	"googlemaps.github.io/maps"
)

func main() {
	// API Key setup
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	apiKey := os.Getenv("GOOGLE_MAPS_API_KEY") // Consider using environment variables
	if apiKey == "" {
		log.Fatalf("GOOGLE_MAPS_API_KEY environment variable is not set")
	}
	client, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("fatal error creating client: %s", err)
	}

	// Get search query (you might fetch this from command-line arguments later)
	query := "Board Game Stores"

	// Call your mapsapi function
	results, err := mapsapi.GetSearchResults(context.Background(), client, query)
	if err != nil {
		log.Printf("Error fetching search results: %s", err)
		return // Or exit with an error code
	}
	fmt.Printf("Results: %v", results)
}
