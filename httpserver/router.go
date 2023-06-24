package httpserver

import (
	"github.com/whatattitude/go-kit-base/httpserver/controller"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/hello", controller.Hello)
}
