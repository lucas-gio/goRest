package goRest

import (
	"github.com/gin-gonic/gin"
)

type IController interface {
	InitializeRoutes(*gin.Engine)
}
