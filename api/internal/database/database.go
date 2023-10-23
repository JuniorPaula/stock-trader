package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI("mongodb://admin:admin@localhost:27017")
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

	db := client.Database("stocktrader")

	return db, nil
}
