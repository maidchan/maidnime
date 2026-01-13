package main

import (
	"backend/internal/api"
	"log"
)

func main() {
	r := api.SetupRouter()

	log.Println("Starting backend server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
