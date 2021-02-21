package goRest

import (
	gin "github.com/gin-gonic/gin"
	"net/http"
)

type BicyclesController struct{}

// Implementation of Controller interface method.
func (this *BicyclesController) initializeRoutes(router *gin.Engine) {
	router.GET("/bicycles", this.Bicycles)
}

func (b *BicyclesController) Bicycles(c *gin.Context) {
	/*var datasource DatasourceService = new(DatasourceService).GetInstance()
	var db *mongo.Database = datasource.GetDb()

	db.Collection("bicycles").Find(datasource.GetContext())*/

	//c.HTML(http.StatusOK, gin.H{"title": "Home Page"})
	c.HTML(http.StatusOK, "404.html", gin.H{"title": "Home Page"})

}
