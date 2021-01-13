package route

import (
	"github.com/gin-gonic/gin"
	"go-template/pkg/demo"
)


func LoadRoutes() {

	r := gin.Default()

	// router
	demo.Route(r)

	gin.SetMode(gin.ReleaseMode)

	r.Run(":9000")
}
