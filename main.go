package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	// Define the API route
	router.HandleFunc("/api/getCurrentValues/{station_name}", getCurrentValues).Methods("GET")

	router.HandleFunc("/api/getHourlyAvg/{station_name}/{dt_from_string}/{dt_to_string}", getHourlyAvg).Methods("GET")

	// Start the server
	port := 8080
	log.Printf("Server listening on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
