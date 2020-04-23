package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"stillUp/api"
)


func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("API_PORT")

	api.Start(port)
}
