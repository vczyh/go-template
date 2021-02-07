package route

import (
	"context"
	"errors"
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
		}
		handlerFunc(nil, c)
	}
}

// server
type HttpServer struct {
	ctx context.Context

	addr        string
	middlewares []gin.HandlerFunc
	routes      []Route
}

func NewHttpServer(ctx context.Context, addr string) *HttpServer {
	return &HttpServer{
		ctx:  ctx,
		addr: addr,
	}
}

func (h *HttpServer) Bind(addr string) {
	h.addr = addr
}

func (h *HttpServer) AddMiddleware(middle gin.HandlerFunc) {
	h.middlewares = append(h.middlewares, middle)
}

func (h *HttpServer) AddRoute(routes ...Route) {
	h.routes = append(h.routes, routes...)
}

func (h *HttpServer) Run(mode string) (err error) {
	defer func() {
		if p := recover(); p != nil {
			err = errors.New(p.(string))
		}
	}()
	// PANIC
	gin.SetMode(mode)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.Use(h.middlewares...)
	for _, route := range h.routes {
		route(r)
	}
	if err = r.Run(h.addr); err != nil {
		return err
	}
	return nil
}

func (h *HttpServer) WithContextKeyAndValueMiddle(key string, valFunc ValueFunc) gin.HandlerFunc {
	return ContextKeyAndValueMiddle(h.ctx, key, valFunc)
}
