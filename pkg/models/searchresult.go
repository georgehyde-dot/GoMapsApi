package models

type SearchResult struct {
	Id          string
	Name        string
	Address     string
	PhoneNumber string
	Website     string
}

func (sr SearchResult) SaveSearchResult() error {

	return nil
}
