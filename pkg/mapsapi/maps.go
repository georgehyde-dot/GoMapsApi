package mapsapi

import (
	"context"

	"github.com/georgehyde-dot/GoMapsApi/pkg/models" // Make sure the import path is correct
	"googlemaps.github.io/maps"
)

func GetSearchResults(ctx context.Context, client *maps.Client, query string) (map[int]models.SearchResult, error) {
	// ... your Google Maps logic will live here
	searchresults := map[int]models.SearchResult{
		1: {Id: 1},
	}

	for result := range searchresults {
		result++
	}

	return searchresults, nil // Or an error
}
