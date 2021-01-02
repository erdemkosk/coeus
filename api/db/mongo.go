package mongo

import (
	"context"
	"sync"

	plugin "github.com/erdemkosk/go-config-service/api/plugin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientInstance *mongo.Client
var clientInstanceError error
var mongoOnce sync.Once

var (
	CONNECTIONSTRING = plugin.GetEnvConfig("MONGO_DB_CONNECTION_STRING")
	DB               = plugin.GetEnvConfig("MONGO_DB_NAME")
	COLLECTION       = "config"
)

func GetMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {

		clientOptions := options.Client().ApplyURI(CONNECTIONSTRING)
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			clientInstanceError = err
		}

		err = client.Ping(context.TODO(), nil)
		if err != nil {
			clientInstanceError = err
		}
		clientInstance = client
	})
	return clientInstance, clientInstanceError
}
