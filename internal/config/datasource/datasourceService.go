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

func (d *DatasourceService) Db() *mongo.Database {
	return d.MongoClient().Database("teste")
}

func (d *DatasourceService) Context() *context.Context {
	return &d.datasource.ctx
}

func (d *DatasourceService) MongoClient() *mongo.Client {
	d.onceDatasource.Do(func() {
		d.datasource.ctx, d.datasource.cancel = context.WithTimeout(context.Background(), 10*time.Second)

		var connectionString string = d.ConfigurationService.Config().Database.ConnectionString
		var clientOptions *options.ClientOptions = options.Client().ApplyURI(connectionString)

		var err error
		d.datasource.mongoClient, err = mongo.Connect(d.datasource.ctx, clientOptions)
		if err != nil {
			log.Fatal(err)
		}
	})

	return d.datasource.mongoClient
}

func (d *DatasourceService) Disconnect() error {
	(d.datasource.cancel)()
	return d.datasource.mongoClient.Disconnect(d.datasource.ctx)
}

func (d *DatasourceService) MakePingToEngineDB() {
	var err error = d.datasource.mongoClient.Ping(d.datasource.ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
}
