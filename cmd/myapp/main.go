package main

import (
	"flag"
	"fmt"
	"go-template/pkg/env"
	"go-template/pkg/log"
	"go-template/pkg/route"
)

var active = flag.String("active", "", "active profile")

func main() {
	flag.Parse()

	// env
	activeProfile := ".env"
	if *active != "" {
		activeProfile = fmt.Sprintf(".env-%s", *active)
	}
	fmt.Printf("active profile: %s\n", activeProfile)
	if err := env.LoadEnvFile(activeProfile); err != nil {
		panic(err)
	}

	// log
	log.ConfigLog()

	// route
	route.ConfigRoutes()
}
