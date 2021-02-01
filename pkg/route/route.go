package route

import (
	"github.com/gin-gonic/gin"
	"go-template/pkg/demo"
)


func LoadRoutes() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// router
	demo.Route(r)


	r.Run(":9000")
}
