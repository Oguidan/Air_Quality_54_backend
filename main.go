// main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Define your API routes
	router.HandleFunc("/api/airqino-data", getAirqinoData).Methods("GET")

	// Set up other routes as needed

	// Start the server
	port := 8080
	log.Printf("Server listening on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
