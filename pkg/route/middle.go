package route

import (
	"context"
	"github.com/gin-gonic/gin"
)

const contextKey = "context"

func ContextMiddle(ctx context.Context, key string, valFunc func() interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		value := valFunc()
		if val, ok := c.Get(contextKey); ok {
			if ctxVal, ok := val.(context.Context); ok {
				c.Set(contextKey, context.WithValue(ctxVal, key, value))
				return
			}
			// todo  error
		}
		c.Set(contextKey, context.WithValue(ctx, key, value))
	}
}

//func AddModuleName(ctx context.Context, name string) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		if val, ok := c.Get(contextKey); ok {
//			if ctxVal, ok := val.(context.Context); ok {
//				c.Set(contextKey, context.WithValue(ctxVal, "name", name))
//				return
//			}
//			// todo  error
//		}
//		c.Set(contextKey, context.WithValue(ctx, "name", name))
//	}
//}

//func auth(c *gin.Context)  {
//	log.Debug("start auth")
//	c.Next()
//	log.Debug("end auth")
//}
