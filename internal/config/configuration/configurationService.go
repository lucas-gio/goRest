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

func (this *ConfigurationService) LoadConfigurations() {
	this.onceConfiguration.Do(func() {
		var err error
		const configFilepath string = "./internal/config/configuration/config.cfg"
		var configFilePath string = filepath.FromSlash(configFilepath)

		err = config.FromEnv().From(configFilePath).To(&this.configuration)

		if err != nil {
			var errorMsg string = "ConfigurationService file not found in " + configFilePath
			log.Error(errorMsg)
			panic(errorMsg)
		}
	})
}

func (this *ConfigurationService) GetConfig() *Configuration {
	return &this.configuration
}
