package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-template/pkg/demo"
	"os"
)

func ConfigRoutes() {
	fmt.Println(os.Getenv(gin.EnvGinMode))
	//gin.SetMode(os.Getenv("GIN_MODE"))
	r := gin.Default()
	loadRoutes(r)
	r.Run(":8080")
}

func loadRoutes(r *gin.Engine) {
	demo.Route(r)
}
