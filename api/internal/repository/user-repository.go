package repository

import (
	"context"
	"stocktrader/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	MongoDB *mongo.Database
}

// Create is a method to create a new user in the database
// It receives a user model and returns the inserted id and an error
func (r *User) Create(user models.User) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := r.MongoDB.Collection("users").InsertOne(ctx, user)
	if err != nil {
		return primitive.NilObjectID, err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		return oid, nil
	}

	return primitive.NilObjectID, nil
}
