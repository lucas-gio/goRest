package routes

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

	//router.GET("/bicycles", bicyclesController.Bicycles)

	router.GET("/cart", func(c *gin.Context) {
		c.HTML(http.StatusOK, "cart.html", gin.H{"title": "Home Page"})
	})

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"titles": "Home Page"})
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
