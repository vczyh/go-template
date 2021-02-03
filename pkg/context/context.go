package context

import (
	appContext "context"
	"go-template/pkg/env"
	"go-template/pkg/flag"
	"go-template/pkg/log"
)

var Ctx appContext.Context

func CreateContext() {
	ctx, cancel := appContext.WithCancel(appContext.Background())
	Ctx = ctx
	defer cancel()

	flag.ConfigFlag()

	env.MustConfigEnv()

	log.ConfigLog()
}
