package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-template/pkg/demo"
	"go-template/pkg/env"
	"go-template/pkg/log"
	"go-template/pkg/route"
	"io"
	"math/rand"
	"os"
)

const (
	requestIdKey = "requestId"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var active string
	flag.StringVar(&active, "active", "", "active profile")
	flag.Parse()

	// config
	activeProfile := ".env"
	if active != "" {
		activeProfile = fmt.Sprintf(".env-%s", active)
	}
	fmt.Printf("active profile: %s\n", activeProfile)
	if err := env.ConfigEnv(activeProfile); err != nil {
		panic(err)
	}

	// log
	logPath := os.Getenv("LOG_PATH")
	ginLogPath := os.Getenv("GIN_LOG_PATH")
	appFileWriter := log.NewFileWriter(logPath, 10, 5, 30)
	ginFileWriter := log.NewFileWriter(ginLogPath, 10, 5, 30)

	gin.DefaultWriter = io.MultiWriter(ginFileWriter)
	gin.DefaultErrorWriter = io.MultiWriter(ginFileWriter, os.Stdout)

	demoLogger := log.NewLogger("Demo", []string{requestIdKey}, appFileWriter, os.Stdout)
	demo.WithLogger(demoLogger)

	// http server
	s := route.NewHttpServer(ctx, ":8081")

	s.WithContextKeyAndValueMiddle(requestIdKey, func() interface{} {
		return fmt.Sprintf("%d", rand.Int())
	})

	s.AddRoute(demo.Route)

	if err := s.Run(os.Getenv("GIN_MODE")); err != nil {
		panic(err)
	}
}
