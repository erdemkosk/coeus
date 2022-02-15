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
	Id        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	Key       string             `bson:"key"`
	Type      string             `bson:"type"`
	Value     interface{}        `bson:"value"`
}

type ConfigFormatted struct {
    Key       string             `bson:"key"`
	Type      string             `bson:"type"`
	Value     interface{}        `bson:"value"`
}
