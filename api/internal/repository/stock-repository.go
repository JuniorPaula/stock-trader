package repository

import (
	"context"
	"stocktrader/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Stock struct {
	MongoDB *mongo.Database
}

func (s *Stock) Create(stock models.Stock) (primitive.ObjectID, error) {
	result, err := s.MongoDB.Collection("stocks").InsertOne(context.Background(), stock)
	if err != nil {
		return primitive.NilObjectID, err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		return oid, nil
	}

	return primitive.NilObjectID, nil
}
