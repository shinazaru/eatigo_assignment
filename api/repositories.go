package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertIntoDB(collection *mongo.Collection, shortlyData *ShortlyData) error {
	_, err := collection.InsertOne(context.Background(), shortlyData)
	return err
}

func FindShortlyData(shortUUID string, ctx context.Context, collection *mongo.Collection) (ShortlyData, error) {
	filter := bson.M{
		"shortUUID": shortUUID,
	}
	mongoData := collection.FindOne(ctx, filter)
	var shortlyData ShortlyData
	if err := mongoData.Decode(&shortlyData); err != nil {
		return shortlyData, err
	}
	return shortlyData, nil
}

func DeleteShortlyData(shortUUID string, ctx context.Context, collection *mongo.Collection) error {
	filter := bson.M{
		"shortUUID": shortUUID,
	}
	mongoData := collection.FindOne(ctx, filter)
	var shortlyData ShortlyData
	if err := mongoData.Decode(&shortlyData); err != nil {
		return err
	}
	id, _ := primitive.ObjectIDFromHex(shortlyData.ID)
	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})

	return err
}
