module github.com/lucas-gio/goRest/cmd

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/gookit/config/v2 v2.0.21
	github.com/lucas-gio/goRest/configs v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.7.0
)

replace github.com/lucas-gio/goRest/configs => ./../configs
