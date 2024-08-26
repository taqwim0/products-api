package main

import (
	"log"
	"net/http"
	"toggle-features-api/utils"
	"toggle-features-api/view"

	"github.com/gorilla/handlers"
)

func main() {
	utils.InitDB()
	defer utils.DB.Close()

	router := view.RegisterRoutes()

	// CORS middleware
	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),                   // Frontend origin
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}), // Allowed methods
		handlers.AllowedHeaders([]string{"Authorization", "Content-Type"}),           // Allowed headers
		handlers.AllowCredentials(),
	)

	// Wrap the router with the CORS middleware
	server := &http.Server{
		Addr:    ":8080",
		Handler: corsMiddleware(router),
	}

	log.Println("Server running on http://localhost:8080")
	log.Fatal(server.ListenAndServe())
}
