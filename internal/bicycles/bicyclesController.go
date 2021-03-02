package goRest

import (
	"github.com/gin-gonic/gin"
	. "github.com/lucas-gio/goRest/internal/interfaces"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type BicyclesController struct {
	Datasource IDatasourceService `di.inject:"datasourceService"`
}

func (b *BicyclesController) Bicycles(c *gin.Context) {
	var _ *mongo.Database = b.Datasource.Db()
	//*db.Collection("bicycles").Find()

	//c.HTML(http.StatusOK, gin.H{"title": "Home Page"})
	c.HTML(http.StatusOK, "404.html", gin.H{"title": "Home Page"})

}

func (b *BicyclesController) Home(c *gin.Context) {
	c.HTML(http.StatusOK, "bicycles/index.tmpl", "")
}
