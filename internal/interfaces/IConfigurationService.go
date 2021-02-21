package goRest

import . "github.com/lucas-gio/goRest/internal/config/configuration"

type IConfigurationService interface {
	LoadConfigurations()
	GetConfig() *Configuration
}
