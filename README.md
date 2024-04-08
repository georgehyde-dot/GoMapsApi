# GoMapsApi

A Go-based API for finding Board Game Stores using the Google Maps API.

## Prerequisites 

* A Google Maps API Key with the following APIs enabled:
    * Places API
    * Places API (NEW)
* Docker installed

## Building the Docker Image

1. Make sure you have a `Dockerfile` in the project's root directory.
2. Navigate to the project directory and run the build command:

   ```bash
   docker build -t gomapsapi .
