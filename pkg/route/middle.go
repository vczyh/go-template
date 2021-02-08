package route

import (
	"context"
	"github.com/gin-gonic/gin"
)

const contextKey = "context"

type ValueFunc func() interface{}

func ContextKeyAndValueMiddle(ctx context.Context, key string, valFunc ValueFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		value := valFunc()
		if val, ok := c.Get(contextKey); ok {
			if ctxVal, ok := val.(context.Context); ok {
				c.Set(contextKey, context.WithValue(ctxVal, key, value))
				return
			}
		}
		c.Set(contextKey, context.WithValue(ctx, key, value))
	}
}

func AddContextIfNotExist(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		if val, ok := c.Get(contextKey); ok && val != nil {
			return
		}
		c.Set(contextKey,context.WithValue(ctx,"key","vale"))
	}
}
