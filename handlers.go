package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	// "io/ioutil"
	"log"
	"net/http"
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

	// Ensure the response body is closed when done
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

func getHourlyAvg(w http.ResponseWriter, r *http.Request) {
	// Extract station_name, dt_from_string and dt_to_string from the request path
	vars := mux.Vars(r)
	stationName := vars["station_name"]
	dtFromString := vars["dt_from_string"]
	dtToString := vars["dt_to_string"]

	// Construct the API URL
	apiURL := fmt.Sprintf("https://airqino-api.magentalab.it/getHourlyAvg/%s/%s/%s?pivot=true", stationName, dtFromString, dtToString)

	// Make the GET request to the external API
	response, err := http.Get(apiURL)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error decoding API response", http.StatusInternalServerError)
		return
	}

	// Ensure the response body is closed when done
	defer response.Body.Close()

	// Check if the request was successful (status code 200)
	if response.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with status code: %d", response.StatusCode)
	}

	// Read the response body
	// body, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// log.Fatal(err)
	// }

	// Print or process the response body (text/csv in this case)
	// fmt.Println(string(body))
	fmt.Print(response)
}
