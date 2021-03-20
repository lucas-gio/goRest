package goRest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PartsController struct {
	partsService iPartsService `di.inject:"partsService"`
}

func (p *PartsController) Parts(c *gin.Context) {
	var max string = c.DefaultQuery("max", "25")
	var offset string = c.DefaultQuery("offset", "0")

	var partsList *[]Part = p.partsService.ListParts(max, offset)
	c.JSON(http.StatusOK, gin.H{"parts": *partsList})
}
