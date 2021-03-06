package goRest

import "github.com/gin-gonic/gin"

func (b *BicyclesController) InitializeRoutes(engine *gin.Engine) {
	//engine.GET("/", b.Home)
	engine.GET("/bicycles", b.Bicycles)
}
