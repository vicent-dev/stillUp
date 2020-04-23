package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"stillUp/api"
	"stillUp/redis"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("API_PORT")

	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	redis.Connect(redisHost, redisPort, redisPassword)
	api.Start(port)
}
