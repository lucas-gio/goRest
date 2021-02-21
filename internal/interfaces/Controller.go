package goRest

import "github.com/gin-gonic/gin"

type Controller interface {
	initializeRoutes(*gin.Engine)
}
