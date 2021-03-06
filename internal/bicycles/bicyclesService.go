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

type BicyclesService struct {
	datasource goRest.IDatasourceService `di.inject:"datasourceService"`
}

/*listBicycles returns an slice of bicycles that match the criteria to find. */
func (b *BicyclesService) ListBicycles(max, offset string) (bicyclesList *[]Bicycle) {
	ctx, cancelFunction := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunction()

	var dbHelper IDatabaseHelper = di.GetInstance("databaseHelper").(IDatabaseHelper)
	maxInt := dbHelper.ProcessMax(max)
	offsetInt := dbHelper.ProcessOffset(offset)

	var mongoOptions *options.FindOptions = &options.FindOptions{
		Limit: &maxInt,
		Skip:  &offsetInt,
	}

	log.Info("Searching for bikes in db.")

	var cursor *mongo.Cursor
	var err error
	cursor, err = b.datasource.Db().Collection("bicycles").Find(ctx, bson.M{}, mongoOptions)

	if err != nil {
		log.Error("Error ocurred in find bicycles from database. ", err)
		panic(err)
	}

	bicyclesList = new([]Bicycle)
	if err = cursor.All(ctx, bicyclesList); err != nil {
		log.Error("Error ocurred in fetch bicycles from database. ", err)
		panic(err)
	}

	log.Info("Returning " + fmt.Sprint(len(*bicyclesList)) + " bikes obtained.")
	return
}
