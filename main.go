package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	// Define the API route
	router.HandleFunc("/api/getCurrentValues/{station_name}", getCurrentValues).Methods("GET")

	router.HandleFunc("/api/getHourlyAvg/{station_name}/{dt_from_string}/{dt_to_string}", getHourlyAvg).Methods("GET")

	router.HandleFunc("/api/getRange/{station_name}/{dt_from_string}/{dt_to_string}", getRange).Methods("GET")

	router.HandleFunc("/api/getSessionInfo/{project_name}", getSessionInfo).Methods("GET")

	router.HandleFunc("/api/getSingleDay/{station_name}/{dt_from_string}", getSingleDay).Methods("GET")

	router.HandleFunc("/api/getStationStatus/{station_id}", getStationStatus).Methods("GET")

	router.HandleFunc("/api/getStations/{project_name}", getStations).Methods("GET")

	router.HandleFunc("/api/getStationHourlyAvg/{station_id}", getStationHourlyAvg).Methods("GET")

	// Start the server
	handler := cors.Default().Handler(router)
	port := 8080
	log.Printf("Server listening on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), handler))
}
