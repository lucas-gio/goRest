package goRest

import (
	"github.com/JeremyLoy/config"
	log "github.com/sirupsen/logrus"
	"sync"
)

type ConfigurationService struct {
	onceConfiguration sync.Once
	configuration     Configuration
}

func (c *ConfigurationService) LoadConfigurations() {
	c.onceConfiguration.Do(func() {
		var err error = config.FromEnv().To(&c.configuration)

		if err != nil {
			var errorMsg string = "Error found when load configurations: "
			log.Error(errorMsg, err)
			panic(err)
		}
	})
}

func (c *ConfigurationService) Config() *Configuration {
	return &c.configuration
}
