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

// Create creates a new stock
func (s *Stock) Create(stock models.Stock) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := s.MongoDB.Collection("stocks").InsertOne(ctx, stock)
	if err != nil {
		return primitive.NilObjectID, err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		return oid, nil
	}

	return primitive.NilObjectID, nil
}

// List returns all stocks
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

// Update updates a stock
func (s *Stock) Update(stock models.Stock) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.MongoDB.Collection("stocks").UpdateOne(ctx, bson.M{"_id": stock.ID}, bson.D{{Key: "$set", Value: stock}})
	if err != nil {
		return err
	}

	return nil
}

// Delete deletes a stock
func (s *Stock) Delete(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.MongoDB.Collection("stocks").DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}

// GetByID returns a stock by id
func (s *Stock) GetByID(id primitive.ObjectID) (models.Stock, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var stock models.Stock
	err := s.MongoDB.Collection("stocks").FindOne(ctx, bson.M{"_id": id}).Decode(&stock)
	if err != nil {
		return models.Stock{}, err
	}

	return stock, nil
}
