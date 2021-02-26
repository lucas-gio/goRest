package main

import (
	"github.com/gin-gonic/gin"
	. "github.com/lucas-gio/goRest/internal/config/configuration"
	. "github.com/lucas-gio/goRest/internal/config/logger"
	. "github.com/lucas-gio/goRest/internal/interfaces"
)

func main() {
	new(Logger).InitLog()

	var configurationService IConfigurationService = DefaultContainer.GetconfigurationService()
	configurationService.LoadConfigurations()
	var configuration *Configuration = configurationService.Config()

	var datasource IDatasourceService = DefaultContainer.GetdatasourceService()
	datasource.MongoClient()
	defer datasource.Disconnect()

	router := gin.Default()
	router.Static("/web", "./web")
	router.LoadHTMLGlob("web/templates/*")

	//routes.InitializeRoutes(router)
	//router.GET("/bicycles", b.Bicycles)
	router.GET("/ping2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	_ = router.Run(configuration.Server.Host + ":" + configuration.Server.Port)
}
