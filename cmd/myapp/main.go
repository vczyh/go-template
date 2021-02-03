package main

import (
	"context"
	"go-template/pkg/demo"
	"go-template/pkg/env"
	"go-template/pkg/flag"
	"go-template/pkg/log"
	"go-template/pkg/route"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	flag.ConfigFlag()

	env.MustConfigEnv()

	log.ConfigLog()

	route.AddRoutes(demo.Route)
	// blocking
	route.Run(ctx)
}
