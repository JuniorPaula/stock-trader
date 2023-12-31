package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name       string             `json:"name,omitempty" bson:"name,omitempty"`
	Email      string             `json:"email,omitempty" bson:"email,omitempty"`
	Password   string             `json:"password,omitempty" bson:"password,omitempty"`
	Founds     float64            `json:"founds,omitempty" bson:"founds,omitempty"`
	Portfolios []Portfolio        `json:"portfolios,omitempty" bson:"portfolios,omitempty"`
	Token      string             `json:"token,omitempty" bson:"token,omitempty"`
}
