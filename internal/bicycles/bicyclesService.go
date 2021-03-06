package goRest

import (
	"context"
	"fmt"
	"github.com/lucas-gio/goRest/internal/config/datasource"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type BicyclesService struct {
	datasource goRest.IDatasourceService `di.inject:"datasourceService"`
}

/*listBicycles returns an slice of bicycles that match the criteria to find. */
func (b *BicyclesService) ListBicycles() (bicyclesList *[]Bicycle) {
	ctx, cancelFunction := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunction()

	var cursor *mongo.Cursor
	var err error

	log.Info("Searching for bikes in db.")
	cursor, err = b.datasource.Db().Collection("bicycles").Find(ctx, bson.M{})

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
