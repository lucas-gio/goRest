package goRest

import (
	"github.com/JeremyLoy/config"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"sync"
)

type ConfigurationService struct {
	onceConfiguration sync.Once
	configuration     Configuration
}

func (c *ConfigurationService) LoadConfigurations() {
	c.onceConfiguration.Do(func() {
		var err error

		configFilepath := os.Getenv("GOREST_CONFIG")
		log.Info("Enviroment variable GOREST_CONFIG has value: " + configFilepath)

		configFilepath = filepath.FromSlash(configFilepath)

		err = config.FromEnv().From(configFilepath).To(&c.configuration)

		if err != nil {
			var errorMsg string = "ConfigurationService file not found in " + configFilepath
			log.Error(errorMsg)
			panic(errorMsg)
		}
	})
}

func (c *ConfigurationService) Config() *Configuration {
	return &c.configuration
}
