package demo

import (
	"blog-y/pkg/common/route"
	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine) {
	r := router.Group("demo/v1")

	r.GET("", route.Handle(TestAPI))
}
