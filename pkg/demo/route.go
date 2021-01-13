package demo

import "github.com/gin-gonic/gin"

func Route(r *gin.Engine) {
	r.GET("/demo", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
