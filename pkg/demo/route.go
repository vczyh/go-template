package demo

import (
	"github.com/gin-gonic/gin"
	"go-template/pkg/route"
)

func Route(router *gin.Engine) {
	r := router.Group("demo/v1")

	r.GET("", route.Handle(test))
}

func test(c *gin.Context) (interface{}, error) {

	q := c.Param("q")

	logger.Debug("test request")

	m := map[string]interface{}{
		"q": q,
	}

	return m, nil
}
