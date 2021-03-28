package demo

import (
	"github.com/gin-gonic/gin"
	"go-template/pkg/route"
)

func Route(router *gin.Engine) {
	r := router.Group("demo/v1")

	r.GET("", route.Handle(TestAPI))
}
