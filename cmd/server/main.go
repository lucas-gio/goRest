package main

import (
	"github.com/gin-gonic/gin"
	"github.com/goioc/di"
	. "github.com/lucas-gio/goRest/internal/bicycles"
	. "github.com/lucas-gio/goRest/internal/config/configuration"
	. "github.com/lucas-gio/goRest/internal/config/datasource"
	. "github.com/lucas-gio/goRest/internal/config/logger"
	. "github.com/lucas-gio/goRest/internal/interfaces"
	. "github.com/lucas-gio/goRest/internal/routers"
	log "github.com/sirupsen/logrus"
	"reflect"
) /*
type Dependency string
const (
	ConfigurationService Dependency = "configurationService"
	DatasourceService Dependency = "datasourceService"
	BicyclesService Dependency = "bicyclesService"
	BicyclesController Dependency = "bicyclesController"
)*/

func main() {
	new(Logger).InitLog()

	initDependencies()

	var configurationService IConfigurationService = di.GetInstance("configurationService").(IConfigurationService)
	configurationService.LoadConfigurations()
	var configuration *Configuration = configurationService.Config()

	var datasource IDatasourceService = di.GetInstance("datasourceService").(IDatasourceService)
	datasource.MongoClient()
	defer datasource.Disconnect()

	router := initRouter()

	if err := router.Run(configuration.Server.Host + ":" + configuration.Server.Port); err != nil {
		log.Error("Error ocurred when initializing router. ", err)
		panic(err)
	}
}

/* initDependencies: Initialize the dependency injection container for retrieve instance later*/
func initDependencies() {
	// System
	_, _ = di.RegisterBean("configurationService", reflect.TypeOf(new(ConfigurationService)))
	_, _ = di.RegisterBean("datasourceService", reflect.TypeOf(new(DatasourceService)))

	// Bicycles
	_, _ = di.RegisterBean("bicyclesService", reflect.TypeOf(new(BicyclesService)))
	_, _ = di.RegisterBean("bicyclesController", reflect.TypeOf(new(BicyclesController)))

	_ = di.InitializeContainer()
}

/* initRouter: returns an initialized gin engine router.
When an new controller is added, then you must add it to IncludeRoutesFrom method in order to register yours new routes.
Note the conversion, all singleton instance obtained from dependency container is converted to *IController.
*/
func initRouter() (router *gin.Engine) {
	var routerManager *RoutesManager = new(RoutesManager)

	routerManager.IncludeRoutesFrom(
		di.GetInstance("bicyclesController").(IController),
	)

	router = routerManager.Init()
	return
}
