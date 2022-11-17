package config

import (
	"context"
	"fmt"
	"os"

	"log"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB() *mongo.Client {
	fmt.Print("Connecting database ..")
	//env uri
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	NODE_ENV := os.Getenv("NODE_ENV")
	CONNECTION_DB := os.Getenv("CONNECTION_DB")
	fmt.Println("NODE_ENV: ", NODE_ENV)

	client, err := mongo.NewClient(options.Client().ApplyURI(CONNECTION_DB))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected database: ", CONNECTION_DB)
	return client
}

//Client instance
var DB *mongo.Client = ConnectMongoDB()

//getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("items").Collection(collectionName)
	return collection
}
