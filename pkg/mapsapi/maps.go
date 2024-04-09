package mapsapi

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/georgehyde-dot/GoMapsApi/pkg/models"
	"github.com/joho/godotenv"
	"googlemaps.github.io/maps"
)

func GetSearchResults(ctx context.Context, query string) ([]models.SearchResult, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	apiKey := os.Getenv("GOOGLE_MAPS_API_KEY")
	if apiKey == "" {
		log.Fatalf("GOOGLE_MAPS_API_KEY environment variable is not set")
	}
	client, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("fatal error creating client: %s", err)
	}

	searchRequest := &maps.TextSearchRequest{
		Query: "Board Game Shops in " + query,
	}
	search, err := client.TextSearch(ctx, searchRequest)

	if err != nil {
		return []models.SearchResult{}, fmt.Errorf("error using textsearch API %s", err)
	}

	var apiResults []models.SearchResult
	for _, result := range search.Results {
		detailsRequest := &maps.PlaceDetailsRequest{PlaceID: result.PlaceID}
		detailsResp, err := client.PlaceDetails(ctx, detailsRequest)
		if err != nil {
			continue
		}
		apiResults = append(apiResults, models.SearchResult{
			Id:          result.PlaceID,
			Name:        result.Name,
			Address:     result.FormattedAddress,
			PhoneNumber: detailsResp.InternationalPhoneNumber,
			Website:     detailsResp.URL,
		})
	}
	return apiResults, nil
}

func buildAPIResponse(results []models.SearchResult) interface{} {
	return struct {
		Results []models.SearchResult `json:"results"`
	}{
		Results: results,
	}
}
