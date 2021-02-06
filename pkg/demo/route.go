package demo

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-template/pkg/route"
)

func Route(router *gin.Engine) {
	r := router.Group("demo/v1")

	r.GET("", route.Handle(test))
}

func test(ctx context.Context, c *gin.Context) {
	logger.Debug(ctx, "test request ID", "key11", "val11")
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
