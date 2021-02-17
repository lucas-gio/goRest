module github.com/lucas-gio/configs/datasource

go 1.15

require (
	github.com/lucas-gio/goRest/configs/configuration v0.0.0 // indirect
	github.com/sirupsen/logrus v1.7.1
	go.mongodb.org/mongo-driver v1.4.6
)

replace github.com/lucas-gio/goRest/configs/configuration => ./../configuration
