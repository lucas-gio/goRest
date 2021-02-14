package configs

import (
	"context"
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

var once sync.Once

func (this *Datasource) GetDb() *mongo.Database {
	return this.GetMongoClient().Database("teste")
}

func (this *Datasource) GetMongoClient() *mongo.Client {
	return this.mongoClient
}

func (this *Datasource) GetContext() *context.Context {
	return &this.ctx
}

func (this *Datasource) GetInstance() *Datasource {
	once.Do(func() {
		var pass string = "91701909Aa"
		var dbName string = "goTest"

		this.ctx, this.cancel = context.WithTimeout(context.Background(), 10*time.Second)

		var uri string = "mongodb+srv://gioialucasf:" + pass + "@gotest.kbkdt.mongodb.net/" + dbName + "?retryWrites=true&w=majority"
		var clientOptions *options.ClientOptions = options.Client().ApplyURI(uri)

		var err error
		this.mongoClient, err = mongo.Connect(this.ctx, clientOptions)
		if err != nil {
			log.Fatal(err)
		}
	})

	return this
}

func (this *Datasource) Disconnect() error {
	(this.cancel)()
	return this.GetMongoClient().Disconnect(this.ctx)
}

func (this *Datasource) MakePingToEngineDB() error {
	return this.GetMongoClient().Ping(this.ctx, readpref.Primary())
}
