package main

import (
	"fmt"
	"log"
	"net/http"

	"student-enrollment/config"
	"student-enrollment/routes"
)

func main() {
	// Connect to database
	config.GetDB()

	// Setup routes
	router := routes.SetupRoutes()

	// Start server
	port := ":8080"
	fmt.Printf("Server is running on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
