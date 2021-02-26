package goRest

type Configuration struct {
	Server struct {
		Host string
		Port string
	}
	Database struct {
		ConnectionString string
		DbName           string
	}
}
