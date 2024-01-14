package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

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

// Define a to represente raw of sensor range data
type Raw struct {
	RawData []SensorRange `json:"raw_data"`
}

// Define a struct to represente sensor range data
type SensorRange struct {
	AUX1 float64 `json:"AUX1"`
	// AUX1_INPUT   string  `json:"AUX1_INPUT"`
	AUX2 float64 `json:"AUX2"`
	// AUX2_INPUT   string  `json:"AUX2_INPUT"`
	AUX3         float64 `json:"AUX3"`
	CO           float64 `json:"CO"`
	ExtT         float64 `json:"extT"`
	IntT         float64 `json:"intT"`
	Lat          float64 `json:"lat"`
	Lon          float64 `json:"lon"`
	NO2          float64 `json:"NO2"`
	O3           float64 `json:"O3"`
	PM10         float64 `json:"PM10"`
	PM25         float64 `json:"PM25"`
	RH           float64 `json:"RH"`
	UTCTimestamp string  `json:"utc_timestamp"`
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
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Parse the CSV data
	reader := csv.NewReader(strings.NewReader(string(body)))
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Convert the CSV data to JSON
	var jsonData []map[string]string
	headers := records[0] // Assumes the first row is headers
	for _, row := range records[1:] {
		record := make(map[string]string)
		for i, value := range row {
			record[headers[i]] = value
		}
		jsonData = append(jsonData, record)
	}

	// Write the JSON data to the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jsonData)

	// Print or process the response body (text/csv in this case)
	fmt.Println(jsonData)
}

func getRange(w http.ResponseWriter, r *http.Request) {
	// Extract station_name, dt_from_string and dt_to_string from the request path
	vars := mux.Vars(r)
	stationName := vars["station_name"]
	dtFromString := vars["dt_from_string"]
	dtToString := vars["dt_to_string"]

	// Construct the API URL
	apiURL := fmt.Sprintf("https://airqino-api.magentalab.it/getRange/%s/%s/%s", stationName, dtFromString, dtToString)

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

	// Decode the response JSON
	var body Raw
	if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
		log.Fatal(err)
		http.Error(w, "Error decoding json API response", http.StatusInternalServerError)
	}

	// Write the response body to the response
	fmt.Print(body)
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}

func getSessionInfo(w http.ResponseWriter, r *http.Request) {
	// Extract project_name from the request path
	vars := mux.Vars(r)
	projectName := vars["project_name"]

	// Construct the API URL
	apiURL := fmt.Sprintf("https://airqino-api.magentalab.it/getSessionInfo/%s", projectName)

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
	var airqinoData any
	if err := json.NewDecoder(response.Body).Decode(&airqinoData); err != nil {
		log.Fatal(err)
		http.Error(w, "Error decoding API response", http.StatusInternalServerError)
		return
	}

	// Return the decoded data as JSON
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(airqinoData)

	defer response.Body.Close()
}

func getSingleDay(w http.ResponseWriter, r *http.Request) {
	// Extract the station_name and dt_from_string
	vars := mux.Vars(r)
	stationName := vars["station_name"]
	dtFromString := vars["dt_from_string"]

	// Construct the API URL
	apiURL := fmt.Sprintf("https://airqino-api.magentalab.it/getSingleDay/%s/%s", stationName, dtFromString)

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
	var airqinoData any
	if err := json.NewDecoder(response.Body).Decode(&airqinoData); err != nil {
		log.Fatal(err)
		http.Error(w, "Error decoding the API response", http.StatusInternalServerError)
		return
	}

	// Return the decoded data as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(airqinoData)
}
func getStationStatus(w http.ResponseWriter, r *http.Request) {
	// Extract station_id from the request path
	vars := mux.Vars(r)
	stationId := vars["station_id"]

	// Construct the API URL
	apiURL := fmt.Sprintf("https://airqino-api.magentalab.it/getStationStatus/%s", stationId)

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
	var airquinoData any
	if err := json.NewDecoder(response.Body).Decode(&airquinoData); err != nil {
		log.Fatal(err)
		http.Error(w, "Error decoding API response", http.StatusInternalServerError)
		return
	}

	// Return the decoded data as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(airquinoData)
}

func getStations(w http.ResponseWriter, r *http.Request) {
	// Extract project_name from the request path
	vars := mux.Vars(r)
	projectName := vars["project_name"]

	// Construct the API URL
	apiURL := fmt.Sprintf("https://airqino-api.magentalab.it/getStations/%s", projectName)

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
	var airqinoData any
	if err := json.NewDecoder(response.Body).Decode(&airqinoData); err != nil {
		log.Fatal(err)
		http.Error(w, "Error decoding API response", http.StatusInternalServerError)
		return
	}
}
