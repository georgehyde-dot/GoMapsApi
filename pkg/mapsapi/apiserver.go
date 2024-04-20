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
			log.Println("Bad Request status")
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

	router.HandleFunc("/search", makeHTTPHandlerFunc(s.handleSearch))
	// router.HandleFunc("/search/{query}", makeHTTPHandlerFunc(s.handleSearchLocation))

	log.Println("JSON API server running on port: ", s.listenAddr)

	fs := http.FileServer(http.Dir("./public"))
	router.PathPrefix("/").Handler(fs)

	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleSearch(w http.ResponseWriter, r *http.Request) error {
	switch {
	case r.Method == "GET":
		return s.handleSearchOptions(w, r)
	case r.Method == "POST":
		return s.handleSearchLocation(w, r)
	default:
		log.Printf("method not allowed %v", r.Method)
		return s.handleBadSearch(w, r)
	}
}

func (s *APIServer) handleBadSearch(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte{})
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
	// vars := mux.Vars(r)
	// query := vars["query"]
	query := r.FormValue("location")
	results, err := GetSearchResults(ctx, query)
	if err != nil {
		return fmt.Errorf("error fetching search results: %w", err)
	}

	apiResponse := buildAPIResponse(results)
	return WriteJSON(w, http.StatusOK, apiResponse)
}
