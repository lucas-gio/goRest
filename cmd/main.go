package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/lucas-gio/goRest/routes"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

import . "github.com/lucas-gio/goRest/configs"

func main() {
	var configuration Configuration = Configuration{
		Server: &Server{
			Host: "localhost",
			Port: "8080",
		},
	}

	(&Logger{}).InitLog()

	var datasource *Datasource = new(Datasource).GetInstance()
	defer datasource.Disconnect()

	var err error = datasource.MakePingToEngineDB()
	if err != nil {
		log.Fatal(err)
	}

	datasource.GetDb().Collection("unaDePruebas").InsertOne(context.Background(), bson.M{"campo1": "abc123"})

	/*var configuration Configuration
	var err error

	configuration, err = loadConfigurationPkg()

	if err != nil {
		panic("The configuration file was not found.")
	}
	*/
	router := gin.Default()
	router.Static("/web", "./web")
	router.LoadHTMLGlob("web/templates/*")

	routes.InitializeRoutes(router)

	_ = router.Run(configuration.Server.Host + ":" + configuration.Server.Port)
}

/*
/** load the configuration system for use the config file
func loadConfigurationPkg() (Configuration, error){
	config.WithOptions(config.ParseEnv)

    config.WithOptions( func(options *config.Options) {
		options.Readonly = true
		options.EnableCache = true
	})

	config.AddDriver(yaml.Driver)

    var configFilePath string =  filepath.FromSlash("./configs/config.yml")
	err := config.LoadFiles(configFilePath)
	if err != nil {
		var errorMsg string = "Configuration file not found in " + configFilePath
	    log.Error(errorMsg)
	    panic(errorMsg)
	}

	var configuration Configuration
	config.BindStruct("", &configuration)

    return configuration, err
}*/
