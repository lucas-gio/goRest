package goRest

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type IDatasourceService interface {
	Db() *mongo.Database
	Context() *context.Context
	MongoClient() *mongo.Client
	Disconnect() error
	MakePingToEngineDB()
}
