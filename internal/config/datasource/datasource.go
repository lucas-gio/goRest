package goRest

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type Datasource struct {
	mongoClient *mongo.Client
	ctx         context.Context
	cancel      context.CancelFunc
}
