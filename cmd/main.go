package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

import . "github.com/lucas-gio/goRest/configs"

func main() {
	var configuration Configuration = Configuration{
		Server: &Server{
			Host: "localhost",
			Port: "8080",
		},
	}
	/*var configuration Configuration
	var err error

	configuration, err = loadConfigurationPkg()

	if err != nil {
		panic("The configuration file was not found.")
	}
	*/
	router := gin.Default()
	router.Static("/web", "./web")
	router.LoadHTMLGlob("web/pages/*")

	router.GET("/index", func(c *gin.Context) {
		// Responde un 200, con la página index, y el modelo.
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Home Page"})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run(configuration.Server.Host + ":" + configuration.Server.Port)
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
