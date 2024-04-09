package mapsapi

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
}

type APIfunc func(http.ResponseWriter, *http.Request) error

type APIError struct {
	Error string
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func makeHTTPHandlerFunc(f APIfunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}
	}
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandlerFunc(s.handleSearch))

	log.Println("JSON API server running on port: ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleSearch(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleSearchOptions(w, r)
	}
	if r.Method == "POST" {
		return s.handleSearchLocation(w, r)
	}

	return fmt.Errorf("method not allowed %v", r.Method)
}

func (s *APIServer) handleSearchOptions(w http.ResponseWriter, r *http.Request) error {
	fmt.Printf("Writer %s\n", w)
	fmt.Printf("request %v\n", r)
	return nil
}

func (s *APIServer) handleSearchLocation(w http.ResponseWriter, r *http.Request) error {
	//todo query extraction & error handling
	ctx := context.Background()
	query := r.Header.Get("location")
	results, err := GetSearchResults(ctx, query)
	if err != nil {
		return fmt.Errorf("error fetching search results: %w", err)
	}

	apiResponse := buildAPIResponse(results)
	return WriteJSON(w, http.StatusOK, apiResponse)
}
