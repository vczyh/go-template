package route

import (
	"context"
	"github.com/gin-gonic/gin"
)

type Route func(r *gin.Engine)

type HandlerFunc func(context.Context, *gin.Context)

func Handle(handlerFunc HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if val, ok := c.Get(contextKey); ok {
			if ctxVal, ok := val.(context.Context); ok {
				handlerFunc(ctxVal, c)
				return
			}
			// todo error
		}
	}
}

//var routes []Route

//func AddRoutes(rs ...Route) {
//	routes = append(routes, rs...)
//}

//func Run(ctx context.Context) {
//	gin.SetMode(os.Getenv("GIN_MODE"))
//	r := gin.Default()
//	r.Use(AddRequestId(ctx))
//	loadRoutes(r)
//	if err := r.Run(":8081"); err != nil {
//		panic(err)
//	}
//}

//func loadRoutes(r *gin.Engine) {
//	for _, route := range routes {
//		route(r)
//	}
//}

// server
type HttpServer struct {
	Ctx context.Context

	addr        string
	middlewares []gin.HandlerFunc
	routes      []Route
}

func NewHttpServer(ctx context.Context, addr string) *HttpServer {
	return &HttpServer{
		Ctx: ctx,
		addr: addr,
	}
}

func (h *HttpServer) Bind(addr string) {
	h.addr = addr
}

func (h *HttpServer) AddMiddleware(middle gin.HandlerFunc) {
	h.middlewares = append(h.middlewares, middle)
}

func (h *HttpServer) AddRoute(route Route) {
	h.routes = append(h.routes, route)
}

func (h *HttpServer) Run() error {
	r := gin.Default()
	r.Use(h.middlewares...)
	for _, route := range h.routes {
		route(r)
	}
	if err := r.Run(h.addr); err != nil {
		return err
	}
	return nil
}
