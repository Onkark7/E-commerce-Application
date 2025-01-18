package main

import (
	"log"
	data "productmicro/Database"
	routes "productmicro/router"
)

func main() {
	err := data.ConnectDB()

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	router := routes.Router()

	router.Run(":8080")
}
