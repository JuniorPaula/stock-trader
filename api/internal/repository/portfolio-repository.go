package repository

import (
	"context"
	"stocktrader/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Portfolio struct {
	MongoDB *mongo.Database
}

// Create is a method to create a new portfolio in the database
// It receives a portfolio model and returns the inserted id and an error
func (r *Portfolio) Create(portfolio models.Portfolio) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := r.MongoDB.Collection("portfolios").InsertOne(ctx, portfolio)
	if err != nil {
		return primitive.NilObjectID, err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		return oid, nil
	}

	return primitive.NilObjectID, nil
}
