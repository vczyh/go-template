package demo

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
)

const contextKey = "context"

func Route(router *gin.Engine) {
	r := router.Group("demo/v1")

	r.GET("", handle(func(ctx context.Context, c *gin.Context) {
		fmt.Println(ctx.Value("reqId"))
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}))

}

type HandlerFunc func(context.Context, *gin.Context)

func handle(handlerFunc HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if val, ok := c.Get(contextKey); ok {
			if ctxVal, ok := val.(context.Context); ok {
				handlerFunc(ctxVal, c)
				return
			}
			// todo error
		}
	}
}
