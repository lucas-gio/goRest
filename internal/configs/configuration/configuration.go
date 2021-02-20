package configuration

import (
	log "github.com/sirupsen/logrus"
	"path/filepath"
	"sync"
)

type Configuration struct {
	Server struct {
		Host string
		Port string
	}
	Database struct {
		ConnectionString string
	}
}

var onceConfiguration sync.Once

const configFilepath string = "./configs/config.cfg"

func (configuration *Configuration) LoadConfigurations() {
	onceConfiguration.Do(func() {
		var err error
		var configFilePath string = filepath.FromSlash(configFilepath)

		err = config.FromEnv().From(configFilePath).To(configuration)

		if err != nil {
			var errorMsg string = "Configuration file not found in " + configFilePath
			log.Error(errorMsg)
			panic(errorMsg)
		}
	})
}
