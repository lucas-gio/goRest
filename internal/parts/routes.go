package goRest

import "github.com/gin-gonic/gin"

func (p *PartsController) InitializeRoutes(engine *gin.Engine) {
	engine.GET("/parts", p.Parts)
}
