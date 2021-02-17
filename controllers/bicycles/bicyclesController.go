package controllers

import (
	gin "github.com/gin-gonic/gin"
	. "github.com/lucas-gio/goRest/configs/datasource"
	"go.mongodb.org/mongo-driver/mongo"

	"net/http"
)

type BicyclesController struct{}

func (this *BicyclesController) Bicycles(c *gin.Context) {
	var datasource Datasource = Datasource.GetInstance()
	var db *mongo.Database = datasource.GetDb()

	db.Collection("bicycles").Find(datasource.GetContext())*/

	c.HTML(http.StatusOK, "cart.html", gin.H{"title": "Home Page"})

}