package main

import (
	"log"
	data "productmicro/Database"
)

func main() {
	err := data.ConnectDB()

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

}
