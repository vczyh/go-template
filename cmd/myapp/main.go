package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-template/pkg/config"
	"go-template/pkg/demo"
	"go-template/pkg/log"
	"go-template/pkg/route"
	"io"
	"math/rand"
	"os"
)

const (
	requestIdKey = "requestId"
)

var (
	configFile = "config.yml"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var active string
	flag.StringVar(&active, "active", "", "active profile")
	flag.Parse()

	if err := config.LoadConfig(configFile); err != nil {
		panic(err)
	}

	// log
	app := log.NewRotate(config.C.Log.App.Path, 10, 5, 30)
	access := log.NewRotate(config.C.Log.AccessLog, 10, 5, 30)

	gin.DefaultWriter = io.MultiWriter(access, os.Stdout)
	gin.DefaultErrorWriter = io.MultiWriter(access, os.Stdout)

	demoLogger := log.NewLogger("Demo", []string{requestIdKey}, app, os.Stdout)
	demo.WithLogger(demoLogger)

	// http server
	s := route.NewHttpServer(ctx, fmt.Sprintf(":%d", config.C.Http.Port))
	s.WithContextKeyAndValueMiddle(requestIdKey, func() interface{} {
		return fmt.Sprintf("%d", rand.Int())
	})
	s.AddRoute(demo.Route)
	if err := s.Run(config.C.Http.Mode); err != nil {
		panic(err)
	}
}
