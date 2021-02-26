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
		var err error

		// MUST be same that out folder of docker last stage.
		//var configFilepath string = filepath.FromSlash("/goRest/gorest.conf")

		err = config.FromEnv(). /*.From(configFilepath)*/ To(&c.configuration)

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
