package route

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
)

const contextKey = "context"

func addRequestId(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqId := fmt.Sprintf("_reqId-%d", rand.Int())
		if val, ok := c.Get(contextKey); ok {
			if ctxVal, ok := val.(context.Context); ok {
				c.Set(contextKey, context.WithValue(ctxVal, "reqId", reqId))
				return
			}
			// todo  error
		}
		c.Set(contextKey, context.WithValue(ctx, "reqId", reqId))
	}
}

//func AddModuleName(ctx context.Context, name string) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		if val, ok := c.Get(contextKey); ok {
//			if ctxVal, ok := val.(context.Context); ok {
//				c.Set(contextKey, log.NewContext(ctxVal, name))
//				return
//			}
//			// todo  error
//		}
//		c.Set(contextKey, NewContext(ctx, name))
//	}
//}

//func auth(c *gin.Context)  {
//	log.Debug("start auth")
//	c.Next()
//	log.Debug("end auth")
//}
