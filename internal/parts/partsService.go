package goRest

import (
	"context"
	"fmt"
	"github.com/goioc/di"
	"github.com/lucas-gio/goRest/internal/config/datasource"
	. "github.com/lucas-gio/goRest/internal/helpers"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type PartsService struct {
	datasource goRest.IDatasourceService `di.inject:"datasourceService"`
}

/*ListParts returns an slice of parts that match the criteria to find. */
func (b *PartsService) ListParts(max, offset string) (partsList *[]Part) {
	ctx, cancelFunction := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunction()

	var dbHelper IDatabaseHelper = di.GetInstance("databaseHelper").(IDatabaseHelper)
	maxInt := dbHelper.ProcessMax(max)
	offsetInt := dbHelper.ProcessOffset(offset)

	var mongoOptions *options.FindOptions = &options.FindOptions{
		Limit: &maxInt,
		Skip:  &offsetInt,
	}

	log.Info("Searching for parts in db.")
	var cursor *mongo.Cursor
	var err error
	cursor, err = b.datasource.Db().Collection("parts").Find(ctx, bson.M{}, mongoOptions)

	if err != nil {
		log.Error("Error ocurred in find parts from database. ", err)
		panic(err)
	}

	partsList = new([]Part)
	if err = cursor.All(ctx, partsList); err != nil {
		log.Error("Error ocurred in fetch parts from database. ", err)
		panic(err)
	}

	log.Info("Returning " + fmt.Sprint(len(*partsList)) + " parts obtained.")
	return
}
