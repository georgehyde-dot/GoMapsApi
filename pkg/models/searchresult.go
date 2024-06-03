package models

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type SearchResult struct {
	Id          string `json:"placeId"`
	Name        string `json:"placeName"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber"`
	Website     string `json:"websiteUrl"`
}

func (sr SearchResult) SaveSearchResult(q string) error {
	filename := fmt.Sprintf("%s_query.csv", q)
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Could not create CSV file: %s", err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)

	records := []string{
		sr.Id,
		sr.Address,
		sr.Name,
		sr.PhoneNumber,
		sr.Website,
	}

	if err := writer.Write(records); err != nil {
		log.Fatalf("Error writing to CSV: %s", err)
	}
	return nil
}
