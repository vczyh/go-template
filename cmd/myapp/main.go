package main

import (
	"blog-y/pkg/blog"
	"blog-y/pkg/common/config"
	"blog-y/pkg/common/log"
	"blog-y/pkg/common/mysql"
	"blog-y/pkg/common/route"
	"blog-y/pkg/demo"
	"fmt"
	"github.com/spf13/cobra"
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
	// blog
	blogLog := log.New("Blog", c.GetString("log.level"), writer, os.Stdout)
	blog.WithLogger(blogLog)

	appLog.Infof("config file: %s", configFile)

	// MySQL
	db, err := mysql.New(c.GetString("mysql.host"),
		c.GetInt("mysql.port"),
		c.GetString("mysql.user"),
		c.GetString("mysql.password"),
		c.GetString("mysql.dbname"))
	if err != nil {
		panic(err)
	}
	blog.WithMySQL(db)
	appLog.Infof("connect mysql successfully")

	// http server
	appLog.Infof("http initialized with port: %d", c.GetInt("http.port"))
	s := route.NewHttpServer(fmt.Sprintf(":%d", c.GetInt("http.port")))
	s.AccessWriters(httpAccessWriter, os.Stdout)
	s.ErrWriters(httpErrWriter, os.Stdout)
	s.AddRoute(demo.Route, blog.Route)
	if err := s.Run(c.GetString("http.mode")); err != nil {
		panic(err)
	}
}
