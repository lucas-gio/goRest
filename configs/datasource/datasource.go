package datasource

import (
	"context"
	. "github.com/lucas-gio/goRest/configs/configuration"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"sync"
	"time"
)

type Datasource struct {
	mongoClient *mongo.Client
	ctx         context.Context
	cancel      context.CancelFunc
}

var onceDatasource sync.Once

func (this *Datasource) GetDb() *mongo.Database {
	return this.GetMongoClient().Database("teste")
}

func (this *Datasource) GetMongoClient() *mongo.Client {
	return this.mongoClient
}

func (this *Datasource) GetContext() *context.Context {
	return &this.ctx
}

func (this *Datasource) ConnectToMongodb(configuration *Configuration) {
	onceDatasource.Do(func() {
		this.ctx, this.cancel = context.WithTimeout(context.Background(), 10*time.Second)

		var clientOptions *options.ClientOptions = options.Client().ApplyURI(configuration.Database.ConnectionString)

		var err error
		this.mongoClient, err = mongo.Connect(this.ctx, clientOptions)
		if err != nil {
			log.Fatal(err)
		}
	})
}

func (this *Datasource) Disconnect() error {
	(this.cancel)()
	return this.GetMongoClient().Disconnect(this.ctx)
}

func (this *Datasource) MakePingToEngineDB() {
	var err error = this.GetMongoClient().Ping(this.ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
}
