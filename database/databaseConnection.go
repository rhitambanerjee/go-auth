package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBInstance() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file")
	}
	MongoDb := os.Getenv("MONGO_URI")

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(MongoDb))

if err != nil {
    log.Fatal(err)
}

	fmt.Print("Connected to MongoDB")

	return client
}

var Client *mongo.Client = DBInstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = (*mongo.Collection)(client.Database("cluster0").Collection(collectionName))
	return collection
}