package goRest

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type IDatasourceService interface {
	GetDb() *mongo.Database
	GetContext() *context.Context
	GetMongoClient() *mongo.Client
	Disconnect() error
	MakePingToEngineDB()
}
