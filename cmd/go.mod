module github.com/lucas-gio/goRest/cmd

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/lucas-gio/goRest/configs v0.0.0-00010101000000-000000000000
	github.com/lucas-gio/goRest/routes v0.0.0-00010101000000-000000000000
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/stretchr/testify v1.6.1 // indirect
	golang.org/x/sys v0.0.0-20200720211630-cb9d2d5c5666 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace github.com/lucas-gio/goRest/configs => ./../configs

replace github.com/lucas-gio/goRest/routes => ./../routes
