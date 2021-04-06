package main

import "time"

type ShortlyData struct {
	ShortUUID string    `json:"shortUUID" bson:"shortUUID"`
	Uri       string    `json:"uri" bson:"uri"`
	ExpireAt  time.Time `json:"expireAt" bson:"expireAt"`
}

type ShortlyMongoModel struct {
	ID        string    `json:"id" bson:"_id"`
	ShortUUID string    `json:"shortUUID" bson:"shortUUID"`
	Uri       string    `json:"uri" bson:"uri"`
	ExpireAt  time.Time `json:"expireAt" bson:"expireAt"`
}
