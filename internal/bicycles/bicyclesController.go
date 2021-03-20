package goRest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BicyclesController struct {
	bicyclesService IBicyclesService `di.inject:"bicyclesService"`
}

func (b *BicyclesController) Bicycles(c *gin.Context) {
	var max string = c.DefaultQuery("max", "25")
	var offset string = c.DefaultQuery("offset", "0")

	var bicyclesList *[]Bicycle = b.bicyclesService.ListBicycles(max, offset)
	c.JSON(http.StatusOK, gin.H{"bicycles": *bicyclesList})
}

func (b *BicyclesController) Home(c *gin.Context) {
	c.HTML(http.StatusOK, "bicycles/index.tmpl", "")
}
