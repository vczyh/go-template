package route

import (
	"errors"
	"github.com/gin-gonic/gin"
	"io"
)

// server
type HttpServer struct {
	addr          string
	middlewares   []gin.HandlerFunc
	routes        []Route
	accessWriters []io.Writer
	errWriters    []io.Writer
}

type Route func(r *gin.Engine)

func NewHttpServer(addr string) *HttpServer {
	return &HttpServer{
		addr: addr,
	}
}

func (h *HttpServer) Bind(addr string) {
	h.addr = addr
}

func (h *HttpServer) AddMiddleware(middles ...gin.HandlerFunc) {
	h.middlewares = append(h.middlewares, middles...)
}

func (h *HttpServer) AddRoute(routes ...Route) {
	h.routes = append(h.routes, routes...)
}

func (h *HttpServer) AccessWriters(writers ...io.Writer) {
	h.accessWriters = append(h.accessWriters, writers...)
}

func (h *HttpServer) ErrWriters(writers ...io.Writer) {
	h.errWriters = append(h.errWriters, writers...)
}

func (h *HttpServer) Run(mode string) (err error) {
	defer func() {
		if p := recover(); p != nil {
			err = errors.New(p.(string))
		}
	}()
	// PANIC
	gin.SetMode(mode)

	gin.DefaultWriter = io.MultiWriter(h.accessWriters...)
	gin.DefaultErrorWriter = io.MultiWriter(h.errWriters...)

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
