package goRest

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Datasource struct {
	mongoClient *mongo.Client
}
