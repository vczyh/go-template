package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-template/pkg/config"
	"go-template/pkg/demo"
	"go-template/pkg/log"
	"go-template/pkg/route"
	"os"
)

var (
	configFile string

	rootCmd = &cobra.Command{
		Use: "myapp",
	}
)

func main() {
	// flag
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "config.yml", "config file")
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}

	// config
	f, err := config.LoadFile(configFile)
	if err != nil {
		panic(err)
	}
	if err = config.ReadFromMultiSources(&config.Env{}, f); err != nil {
		panic(err)
	}

	// log
	writer := log.NewRotate(config.GetLogPath(), 10, 5, 30)
	appLogger := log.New("", config.GetLogLevel(), writer, os.Stdout)

	appLogger.Infof("Load config file: %s", configFile)
	appLogger.Infof("Config keys and values: %s", config.Info())
	appLogger.Infof("HTTP initialized with port: %d", config.GetHttpPort())

	// http server
	s := route.NewHttpServer(fmt.Sprintf(":%d", config.GetHttpPort()))
	//s.AccessWriters(httpAccessWriter)
	//s.ErrWriters(httpErrWriter, os.Stdout)
	s.AddMiddleware(route.TraceLoggerMiddleware(appLogger))
	s.AddRoute(demo.Route)
	if err := s.Run(config.GetHttpMode()); err != nil {
		panic(err)
	}
}
