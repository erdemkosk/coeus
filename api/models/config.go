package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Person struct {
	ID        uint   `json:”id”`
	FirstName string `json:”firstname”`
	LastName  string `json:”lastname”`
}

type Config struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	Key       string             `json:"key" bson:"key"`
	Type      string             `json:"type" bson:"type"`
	Value     interface{}        `json:"value" bson:"value"`
}

type ConfigFormatted struct {
    Key       string             `bson:"key"`
	Type      string             `bson:"type"`
	Value     interface{}        `bson:"value"`
}
