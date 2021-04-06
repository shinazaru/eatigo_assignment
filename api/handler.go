package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
	"github.com/renstrom/shortuuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateShortlyHandler(c echo.Context) (err error) {
	response := make(map[string]interface{})
	shortlyData := new(ShortlyData)
	if err = c.Bind(shortlyData); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	dbClient, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URL")))
	if err != nil {
		log.Panicln(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err = dbClient.Connect(ctx); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	collection := dbClient.Database(os.Getenv("MONGO_DATABASE")).Collection("shortlyData")

	shortlyData.ShortUUID = shortuuid.New()
	shortlyData.ExpireAt = time.Now().Add(time.Duration(time.Hour * 24 * 30))
	if err = InsertIntoDB(collection, shortlyData); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	response["uri"] = os.Getenv("API_DOMAIN") + shortlyData.ShortUUID
	return c.JSON(200, response)
}

func RedirectShortlyHandler(c echo.Context) error {
	shortUUID := c.Param("id")

	dbClient, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URL")))
	if err != nil {
		log.Panicln(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err = dbClient.Connect(ctx); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	collection := dbClient.Database(os.Getenv("MONGO_DATABASE")).Collection("shortlyData")
	shortlyData, err := FindShortlyData(shortUUID, ctx, collection)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	shortlyData.ExpireAt = time.Now().Add(time.Duration(time.Hour * 24 * 30))
	id, _ := primitive.ObjectIDFromHex(shortlyData.ID)
	_, err = collection.UpdateOne(ctx, bson.M{"_id": id}, bson.D{{"$set", bson.D{{"expireAt", shortlyData.ExpireAt}}}})

	return c.Redirect(301, shortlyData.Uri)
}

func DeleteShortlyHandler(c echo.Context) error {
	shortUUID := c.Param("id")
	dbClient, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URL")))
	if err != nil {
		log.Panicln(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err = dbClient.Connect(ctx); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	collection := dbClient.Database(os.Getenv("MONGO_DATABASE")).Collection("shortlyData")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = DeleteShortlyData(shortUUID, ctx, collection)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.String(200, "DELETED")
}
