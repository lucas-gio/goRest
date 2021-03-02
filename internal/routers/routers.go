package goRest

import (
	"github.com/gin-gonic/gin"
	. "github.com/lucas-gio/goRest/internal/interfaces"
)

type RoutesManager struct {
	controllerList []IController
}

// Register the route configuration of app
func (r *RoutesManager) IncludeRoutesFrom(controllers ...IController) {
	r.controllerList = append(r.controllerList, controllers...)
}

// initialization
func (r *RoutesManager) Init() (router *gin.Engine) {
	router = gin.Default()

	router.Static("/web", "./web")
	router.LoadHTMLGlob("web/templates/**/*")

	for _, controller := range r.controllerList {
		(controller).InitializeRoutes(router)
	}

	return
}

/*
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

router.GET("/404", func(c *gin.Context) {
		c.HTML(http.StatusOK, "404.html", gin.H{"title": "Home Page"})
	})

	router.GET("/accessories", func(c *gin.Context) {
		c.HTML(http.StatusOK, "accessories.html", gin.H{"title": "Home Page"})
	})

	//

	router.GET("/cart", func(c *gin.Context) {
		c.HTML(http.StatusOK, "cart.html", gin.H{"title": "Home Page"})
	})

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{"titles": "Home Page"})
	})

	router.GET("/parts", func(c *gin.Context) {
		c.HTML(http.StatusOK, "parts.html", gin.H{"title": "Home Page"})
	})

	router.GET("/single", func(c *gin.Context) {
		c.HTML(http.StatusOK, "web/pages/single.html", gin.H{"title": "Home Page"})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
*/
