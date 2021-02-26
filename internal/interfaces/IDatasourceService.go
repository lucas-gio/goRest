package goRest

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type IDatasourceService interface {
	Db() *mongo.Database
	MongoClient() *mongo.Client
	Disconnect() error
}
