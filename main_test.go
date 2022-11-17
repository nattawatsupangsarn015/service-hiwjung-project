package main

import (
	"log"
	"os"
	"testing"

	"example/service-hiwjung-project/config"
	"example/service-hiwjung-project/test"

	"github.com/joho/godotenv"
)

func Test(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	NODE_ENV := os.Getenv("NODE_ENV")

	log.Print("connecting db ..")
	config.ConnectMongoDB()
	log.Print("connecting db success !")

	test.TestRouterHomeWithSuccess(t, NODE_ENV)
}
