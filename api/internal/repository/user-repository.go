package repository

import "go.mongodb.org/mongo-driver/mongo"

type User struct {
	MongoDB *mongo.Database
}
