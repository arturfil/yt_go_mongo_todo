package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func ConnectToMongo() (*mongo.Client, error) {
	// MongoDb connection string
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// getting username and password from .env
	username := os.Getenv("MONGO_DB_USERNAME")
	password := os.Getenv("MONGO_DB_PASSWORD")

	// setting auth credentials
	clientOptions.SetAuth(options.Credential{
		Username: username,
		Password: password,
	})

	// Connect to mongo
	// TODO: check either using TODO or Background
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	log.Println("Connected to mongo...")

	return client, nil
}
