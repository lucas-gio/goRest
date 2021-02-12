package configs

type Server struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type Configuration struct {
	Server *Server `mapstructure:"server"`
}
