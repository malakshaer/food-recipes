package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectMongoDB() (*mongo.Client, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return client, nil
}

func OpenCollection(collectionName string) *mongo.Collection {
	client, err := connectMongoDB()
	if err != nil {
		log.Fatal(err)
	}

	var collection *mongo.Collection = client.Database("food-recipes").Collection(collectionName)
	return collection
}
