package route

import (
	"github.com/gin-gonic/gin"
	"go-template/pkg/demo"
)

func ConfigRoutes() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	loadRoutes(r)
	r.Run(":8080")
}

func loadRoutes(r *gin.Engine) {
	demo.Route(r)
}
