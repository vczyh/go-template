package route

import (
	"context"
	"github.com/gin-gonic/gin"
	"os"
)

type Route func(r *gin.Engine)

var routes []Route

func AddRoutes(rs ...Route) {
	routes = append(routes,rs...)
}

func Run(ctx context.Context) {
	gin.SetMode(os.Getenv("GIN_MODE"))
	r := gin.Default()
	r.Use(addRequestId(ctx))
	loadRoutes(r)
	if err := r.Run(":8081"); err != nil {
		panic(err)
	}
}

func loadRoutes(r *gin.Engine) {
	for _, route := range routes {
		route(r)
	}
}
