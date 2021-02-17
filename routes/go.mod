module github.com/lucas-gio/goRest/routes

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/lucas-gio/goRest/controllers v0.0.0-00010101000000-000000000000
)

replace github.com/lucas-gio/goRest/controllers => ./../controllers
