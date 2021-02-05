package main

import (
	"context"
	"fmt"
	"go-template/pkg/demo"
	"go-template/pkg/env"
	"go-template/pkg/flag"
	"go-template/pkg/log"
	"go-template/pkg/route"
	"math/rand"
)

func main() {

	// app context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	flag.ConfigFlag()

	env.MustConfigEnv()

	//log.ConfigLog([]string{"reqId"})



	// log
	log.NewLogger("DEMO",)

	// http
	server := route.NewHttpServer(ctx, ":8081")
	// 为每个请求添加requestID
	server.AddMiddleware(route.ContextMiddle(server.Ctx, "reqId", func() interface{} {
		return fmt.Sprintf("%d", rand.Int())
	}))
	// 添加路由
	server.AddRoute(demo.Route)
	if err := server.Run(); err != nil {
		panic(err)
	}
}
