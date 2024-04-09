package main

import (
	"github.com/georgehyde-dot/GoMapsApi/pkg/mapsapi"
)

func main() {
	server := mapsapi.NewAPIServer(":3000")
	server.Run()
}
