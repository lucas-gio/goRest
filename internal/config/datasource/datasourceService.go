package goRest

import (
	"context"
	"github.com/lucas-gio/goRest/internal/config/configuration"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"sync"
	"time"
)

/*DatasourceService: the service with his dependencies*/
type DatasourceService struct {
	onceDatasource       sync.Once
	datasource           Datasource
	ConfigurationService goRest.IConfigurationService `di.inject:"configurationService"`
}

const timeout time.Duration = 30 * time.Second

/*Db: was created as helper to use in each db query.*/
func (d *DatasourceService) Db() *mongo.Database {
	return d.MongoClient().Database(d.ConfigurationService.Config().Database.DbName)
}

/*MongoClient: instance the singleton Datasource.mongoClient. If cant connect to Db, then an log fatal and panic is generated. */
func (d *DatasourceService) MongoClient() *mongo.Client {
	d.onceDatasource.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		var connectionString string = d.ConfigurationService.Config().Database.ConnectionString

		log.Info("Connecting to " + connectionString)
		var clientOptions *options.ClientOptions = options.Client().ApplyURI(connectionString)

		var err error
		d.datasource.mongoClient, err = mongo.Connect(ctx, clientOptions)

		if err != nil {
			log.Fatal("Mongo connect: cant connect to Db: ", err)
			panic(err)
		}

		ctx, cancel = context.WithTimeout(context.Background(), timeout)
		defer cancel()

		err = d.datasource.mongoClient.Ping(ctx, readpref.Primary())
		if err != nil {
			log.Fatal("Mongo ping: response not received: ", err)
			panic(err)
		}

		log.Info("Mongo ping: received!!!")
		log.Info("Connection to DB success")
	})

	return d.datasource.mongoClient
}

/*Disconnect: make an close of db connection*/
func (d *DatasourceService) Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	log.Info("Mongo disconnect: closing db connection")

	if err := d.datasource.mongoClient.Disconnect(ctx); err != nil {
		log.Error("Mongo disconnect: error detected when disconnect from db. ", err)
		panic(err)
	}
}
