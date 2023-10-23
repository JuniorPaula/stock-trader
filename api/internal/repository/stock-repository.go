package repository

import (
	"context"
	"log"
	"stocktrader/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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

func (s *Stock) List() ([]models.Stock, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := s.MongoDB.Collection("stocks").Find(ctx, bson.D{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var stocks []models.Stock
	err = cursor.All(ctx, &stocks)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		return nil, err
	}

	return stocks, nil
}
