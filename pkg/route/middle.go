package route

import (
	"github.com/gin-gonic/gin"
	"go-template/pkg/context"
	"math/rand"
	"strconv"
)

func addRequestId(c *gin.Context) {
	ctx := NewContext(context.Ctx, strconv.Itoa(rand.Int()))
	c.Set("context", ctx)
}

//func auth(c *gin.Context)  {
//	log.Debug("start auth")
//	c.Next()
//	log.Debug("end auth")
//}
