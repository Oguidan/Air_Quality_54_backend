// handlers.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Define a struct to represent Airqino data
type AirqinoData struct {
	StationName string       `json:"station_name"`
	Timestamp   string       `json:"timestamp"`
	Values      []SensorData `json:"values"`
}

// Define a struct to represent sensor data
type SensorData struct {
	Sensor string  `json:"sensor"`
	Unit   string  `json:"unit"`
	Value  float64 `json:"value"`
}

func getCurrentValues(w http.ResponseWriter, r *http.Request) {
	// Extract station_name from the request path
	vars := mux.Vars(r)
	stationName := vars["station_name"]

	// Construct the API URL
	apiURL := fmt.Sprintf("https://airqino-api.magentalab.it/getCurrentValues/%s", stationName)

	// Make the GET request to the external API
	response, err := http.Get(apiURL)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error making external API request", http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	// Decode the response JSON
	var airqinoData AirqinoData
	if err := json.NewDecoder(response.Body).Decode(&airqinoData); err != nil {
		log.Fatal(err)
		http.Error(w, "Error decoding API response", http.StatusInternalServerError)
		return
	}

	// Return the decoded data as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(airqinoData)
	fmt.Print(airqinoData)
}
