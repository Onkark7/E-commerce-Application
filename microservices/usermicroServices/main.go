package main

import (
	"log"
	"net/http"
	data "user/database"
	"user/routes"
)

func main() {
	err := data.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	router := routes.RegisterRoutes()

	log.Println("Starting server on port 8080...")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
