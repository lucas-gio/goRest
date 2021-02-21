package goRest

import (
	"context"
	. "github.com/lucas-gio/goRest/internal/interfaces"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"sync"
	"time"
)

type DatasourceService struct {
	onceDatasource       sync.Once
	datasource           Datasource
	ConfigurationService IConfigurationService
}

func (this *DatasourceService) GetDb() *mongo.Database {
	return this.GetMongoClient().Database("teste")
}

func (this *DatasourceService) GetContext() *context.Context {
	return &this.datasource.ctx
}

func (this *DatasourceService) GetMongoClient() *mongo.Client {
	this.onceDatasource.Do(func() {
		this.datasource.ctx, this.datasource.cancel = context.WithTimeout(context.Background(), 10*time.Second)

		var clientOptions *options.ClientOptions = options.Client().ApplyURI(this.ConfigurationService.GetConfig().Database.ConnectionString)

		var err error
		this.datasource.mongoClient, err = mongo.Connect(this.datasource.ctx, clientOptions)
		if err != nil {
			log.Fatal(err)
		}
	})

	return this.datasource.mongoClient
}

func (this *DatasourceService) Disconnect() error {
	(this.datasource.cancel)()
	return this.datasource.mongoClient.Disconnect(this.datasource.ctx)
}

func (this *DatasourceService) MakePingToEngineDB() {
	var err error = this.datasource.mongoClient.Ping(this.datasource.ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
}
