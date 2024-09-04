package main

import (
	"log"
	"net/http"
	"toggle-features-api/utils"
	"toggle-features-api/view"
)

func main() {
	utils.InitDB()
	defer utils.DB.Close()

	router := view.RegisterRoutes()

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
