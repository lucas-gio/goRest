package goRest

import "go.mongodb.org/mongo-driver/bson/primitive"

type Part struct {
	Id          primitive.ObjectID `bson:"_id"`
	Code        string             `json:"code"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Origin      string             `json:"origin"`
	UnitPrice   string             `json:"unitPrice"`
	Image       string             `json:"image"`
}
