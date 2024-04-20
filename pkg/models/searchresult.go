package models

type SearchResult struct {
	Id          string `json:"placeId"`
	Name        string `json:"placeName"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber"`
	Website     string `json:"websiteUrl"`
}

func (sr SearchResult) SaveSearchResult() error {

	return nil
}
