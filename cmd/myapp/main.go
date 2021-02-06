package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-template/pkg/demo"
	"go-template/pkg/env"
	"go-template/pkg/flag"
	"go-template/pkg/log"
	"go-template/pkg/route"
	"io"
	"math/rand"
	"os"
)

func main() {

	// app context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// flag
	flag.ConfigFlag()

	// config
	env.MustConfigEnv()



	// log
	logPath := os.Getenv("LOG_PATH")
	appFileWriter := log.NewFileWriter(logPath, 10, 5, 30)
	ginFileWriter := log.NewFileWriter("gin.log", 10, 5, 30)

	gin.DefaultWriter = io.MultiWriter(ginFileWriter, os.Stdout)
	gin.DefaultErrorWriter = io.MultiWriter(ginFileWriter, os.Stdout)

	demoLogger := log.NewLogger("Demo", []string{"reqId"}, appFileWriter, os.Stdout)
	demo.WithLogger(demoLogger)

	// http server
	s := route.NewHttpServer(ctx, ":8081")

	s.AddMiddleware(route.ContextKeyAndValueMiddle(s.Ctx, "reqId", func() interface{} {
		return fmt.Sprintf("%d", rand.Int())
	}))

	s.AddRoute(demo.Route)

	if err := s.Run(); err != nil {
		panic(err)
	}
}
