package main

import (
	"fmt"
	"log"
	"os"

	"example/service-hiwjung-project/api"
	"example/service-hiwjung-project/config"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	NODE_ENV := os.Getenv("NODE_ENV")
	CONNECTION_DB := os.Getenv("CONNECTION_DB")

	fmt.Printf(CONNECTION_DB)

	log.Print("connecting db ..")
	config.ConnectMongoDB()
	log.Print("connecting db success !")

	log.Print("Starting the service .. ")
	router := api.Router(NODE_ENV)
	router.Run()
}
