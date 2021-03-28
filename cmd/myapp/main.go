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

	// config
	c, err := config.New(configFile)
	if err != nil {
		panic(err)
	}
	demo.WithConfig(c)

	// log
	writer := log.NewRotate(c.GetString("log.path"), 10, 5, 30)
	httpAccessWriter := log.NewRotate(c.GetString("log.http.access-log"), 10, 5, 30)
	httpErrWriter := log.NewRotate(c.GetString("log.http.error-log"), 10, 5, 30)
	// app
	appLog := log.New("App", c.GetString("log.level"), writer, os.Stdout)
	// demo
	demoLog := log.New("Demo", c.GetString("log.level"), writer, os.Stdout)
	demo.WithLogger(demoLog)

	appLog.Infof("config file: %s", configFile)
	appLog.Infof("http initialized with port: %d", c.GetInt("http.port"))

	// http server
	s := route.NewHttpServer(fmt.Sprintf(":%d", c.GetInt("http.port")))
	s.AccessWriters(httpAccessWriter)
	s.ErrWriters(httpErrWriter, os.Stdout)
	s.AddRoute(demo.Route)
	if err := s.Run(c.GetString("http.mode")); err != nil {
		panic(err)
	}
}
