module github.com/lucas-gio/goRest/cmd

go 1.15

replace github.com/lucas-gio/goRest/configs/configuration => ./../configs/configuration

replace github.com/lucas-gio/goRest/configs/datasource => ./../configs/datasource

replace github.com/lucas-gio/goRest/configs/logger => ./../configs/logger

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/lucas-gio/goRest/configs/configuration v0.0.0
	github.com/lucas-gio/goRest/configs/datasource v0.0.0
	github.com/lucas-gio/goRest/configs/logger v0.0.0
	github.com/lucas-gio/goRest/routes v0.0.0-20210214025155-939cc15e12a1
	github.com/sirupsen/logrus v1.7.1
	go.mongodb.org/mongo-driver v1.4.6
)
