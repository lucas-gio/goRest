package goRest

import (
	"github.com/JeremyLoy/config"
	log "github.com/sirupsen/logrus"
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
		const configFilepath string = "/home/pc/config.cfg"
		var configFilePath string = filepath.FromSlash(configFilepath)

		err = config.FromEnv().From(configFilePath).To(&c.configuration)

		if err != nil {
			var errorMsg string = "ConfigurationService file not found in " + configFilePath
			log.Error(errorMsg)
			panic(errorMsg)
		}
	})
}

func (c *ConfigurationService) GetConfig() *Configuration {
	return &c.configuration
}
