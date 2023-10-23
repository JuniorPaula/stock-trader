package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	connectString = ""
)

func Connect() (*mongo.Database, error) {
	connectString = fmt.Sprintf("mongodb://%s:%s@%s:%s",
		os.Getenv("MONGO_INITDB_ROOT_USERNAME"),
		os.Getenv("MONGO_INITDB_ROOT_PASSWORD"),
		os.Getenv("MONGO_HOST"),
		os.Getenv("MONGO_PORT"),
	)

	clientOptions := options.Client().ApplyURI(connectString)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	time.AfterFunc(10*time.Second, func() {
		client.Disconnect(context.Background())
	})

	db := client.Database(os.Getenv("MONGO_INITDB_DATABASE"))

	return db, nil
}
