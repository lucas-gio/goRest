module github.com/goRest/controllers/bicycles

go 1.15

replace github.com/lucas-gio/goRest/configs/datasource => ./../../configs/datasource

require (
	github.com/gin-gonic/gin v1.6.3
	go.mongodb.org/mongo-driver v1.4.6
	github.com/lucas-gio/goRest/configs/datasource v0.0.0
)
