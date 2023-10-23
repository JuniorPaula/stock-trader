package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Portfolio struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID   primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	StockID  primitive.ObjectID `json:"stock_id,omitempty" bson:"stock_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	Price    float64            `json:"price,omitempty" bson:"price,omitempty"`
	Quantity int                `json:"quantity,omitempty" bson:"quantity,omitempty"`
}
