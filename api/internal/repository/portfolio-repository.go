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

// GetAllByUserID is a method to get all portfolios by user id
// It receives a user id and returns a slice of portfolios and an error
func (r *Portfolio) GetAllByUserID(userID primitive.ObjectID) ([]models.Portfolio, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var portfolios []models.Portfolio

	cursor, err := r.MongoDB.Collection("portfolios").Find(ctx, map[string]interface{}{"user_id": userID})
	if err != nil {
		return portfolios, err
	}

	err = cursor.All(ctx, &portfolios)
	if err != nil {
		return portfolios, err
	}

	return portfolios, nil
}

// GetByID is a method to get a portfolio by id
// It receives a portfolio id and returns a portfolio model and an error
func (r *Portfolio) GetByID(id primitive.ObjectID) (models.Portfolio, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var portfolio models.Portfolio

	err := r.MongoDB.Collection("portfolios").FindOne(ctx, map[string]interface{}{"_id": id}).Decode(&portfolio)
	if err != nil {
		return portfolio, err
	}

	return portfolio, nil
}

// Update is a method to update a portfolio
// It receives a portfolio model and returns an error
func (r *Portfolio) Update(portfolio models.Portfolio) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.MongoDB.Collection("portfolios").UpdateOne(ctx, map[string]interface{}{"_id": portfolio.ID}, map[string]interface{}{"$set": portfolio})
	if err != nil {
		return err
	}

	return nil
}

// Delete is a method to delete a portfolio
// It receives a portfolio model and returns an error
func (r *Portfolio) Delete(portfolio models.Portfolio) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.MongoDB.Collection("portfolios").DeleteOne(ctx, map[string]interface{}{"_id": portfolio.ID})
	if err != nil {
		return err
	}

	return nil
}
