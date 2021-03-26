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
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "config.yml", "config file")

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}

	c, err := config.New(configFile)
	if err != nil {
		panic(err)
	}

	//if err := config.LoadConfig(configFile); err != nil {
	//	panic(err)
	//}

	// log
	appWriter := log.NewRotate(c.GetString("log.path"), 10, 5, 30)
	accessWriter := log.NewRotate(c.GetString("log.http.access-log"), 10, 5, 30)
	errWriter := log.NewRotate(c.GetString("log.http.error-log"), 10, 5, 30)

	appLogger := log.New("App", c.GetString("log.level"), appWriter, os.Stdout)
	demoLogger := log.New("Demo", c.GetString("log.level"), appWriter, os.Stdout)
	demo.WithLogger(demoLogger)

	appLogger.Infof("config file: %s", configFile)
	appLogger.Infof("http initialized with port: %d", c.Get("http.port"))

	// http server
	s := route.NewHttpServer(fmt.Sprintf(":%d", c.GetInt("http.port")))
	s.AccessWriters(accessWriter)
	s.ErrWriters(errWriter, os.Stdout)
	s.AddRoute(demo.Route)
	if err := s.Run(c.GetString("http.mode")); err != nil {
		panic(err)
	}
}
