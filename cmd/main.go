package main

import (
	"github.com/gin-gonic/gin"
	. "github.com/lucas-gio/goRest/configs/configuration"
	. "github.com/lucas-gio/goRest/configs/datasource"
	. "github.com/lucas-gio/goRest/configs/logger"
	"github.com/lucas-gio/goRest/routes"
)

func main() {
	new(Logger).InitLog()

	var configuration *Configuration = new(Configuration)
	configuration.LoadConfigurations()

	var datasource *Datasource = new(Datasource)
	datasource.ConnectToMongodb(configuration)
	defer datasource.Disconnect()
	datasource.MakePingToEngineDB()

	router := gin.Default()
	router.Static("/web", "./web")
	router.LoadHTMLGlob("web/templates/*")

	routes.InitializeRoutes(router)

	_ = router.Run(configuration.Server.Host + ":" + configuration.Server.Port)
}
