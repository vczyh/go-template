package main

import (
	"flag"
	"fmt"
	"go-template/pkg/config"
	"go-template/pkg/demo"
	"go-template/pkg/log"
	"go-template/pkg/route"
	"os"
)

const (
	requestIdKey = "requestId"
)

var (
	configFile = "config.yml"
)

func main() {
	var active string
	flag.StringVar(&active, "active", "", "active profile")
	flag.Parse()

	if err := config.LoadConfig(configFile); err != nil {
		panic(err)
	}

	// log
	appWriter := log.NewRotate(config.C.Log.App.Path, 10, 5, 30)
	accessWriter := log.NewRotate(config.C.Log.Http.AccessLog, 10, 5, 30)
	errWriter := log.NewRotate(config.C.Log.Http.ErrorLog, 10, 5, 30)

	demoLogger := log.NewLogger("Demo", appWriter, os.Stdout)
	demo.WithLogger(demoLogger)

	// http server
	s := route.NewHttpServer(fmt.Sprintf(":%d", config.C.Http.Port))
	s.AccessWriters(accessWriter)
	s.ErrWriters(errWriter, os.Stdout)
	s.AddRoute(demo.Route)
	if err := s.Run(config.C.Http.Mode); err != nil {
		panic(err)
	}
}
