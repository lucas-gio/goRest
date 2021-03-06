package goRest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BicyclesController struct {
	bicyclesService IBicyclesService `di.inject:"bicyclesService"`
}

func (b *BicyclesController) Bicycles(c *gin.Context) {
	var bicyclesList *[]Bicycle = b.bicyclesService.ListBicycles()
	c.JSON(http.StatusOK, gin.H{"bicycles": *bicyclesList})
}

func (b *BicyclesController) Home(c *gin.Context) {
	c.HTML(http.StatusOK, "bicycles/index.tmpl", "")
}
