package main

import (
	"context"
	"go-template/pkg/env"
	"go-template/pkg/flag"
	"go-template/pkg/log"
	"go-template/pkg/route"
)

var Ctx

func main() {
	//ctx = 1

	// blocking
	route.ConfigRoutes()
}
