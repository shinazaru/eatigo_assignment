package main

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupMongoCollectionIndex() error {
	dbClient, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URL")))
	if err != nil {
		log.Panicln(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err = dbClient.Connect(ctx); err != nil {
		return err
	}
	collection := dbClient.Database(os.Getenv("MONGO_DATABASE")).Collection("shortlyData")
	defer cancel()
	indexModel := mongo.IndexModel{
		Keys: bson.M{
			"expireAt": 1,
		},
		Options: options.Index().SetExpireAfterSeconds(0),
	}
	_, err = collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return err
	}
	return nil
}
