package main

import (
	"flag"
	"fmt"
	"go-template/pkg/config"
	"go-template/pkg/info"
	"go-template/pkg/log"
	"go-template/pkg/route"
)

var active = flag.String("active", "", "active profile")

func main() {

	flag.Parse()

	log.ConfigLog()

	err := info.PrintInfo()
	if err != nil {
		log.Fatal("failed print info")
	}

	activeProfile := ".env"
	if *active != "" {
		activeProfile = fmt.Sprintf(".env-%s", *active)
	}
	log.Info("active profile:", activeProfile)
	if err := config.LoadEnvFile(activeProfile); err != nil {
		log.Fatal("failed to load environment file")
	}

	route.ConfigRoutes()
}
