package goRest

type IConfigurationService interface {
	LoadConfigurations()
	Config() *Configuration
}
