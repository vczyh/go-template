package route

import (
	"github.com/gin-gonic/gin"
	"go-template/pkg/demo"
	"os"
)

func ConfigRoutes() {
	gin.SetMode(os.Getenv("GIN_MODE"))
	r := gin.Default()
	r.Use(addRequestId, auth)
	loadRoutes(r)
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}

func loadRoutes(r *gin.Engine) {
	demo.Route(r)
}
