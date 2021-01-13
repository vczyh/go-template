package main

import (
	"flag"
	"fmt"
	"go-template/pkg/config"
	"go-template/pkg/log"
	"os"
)

func main() {
	// 配置日志
	log.ConfigLog()

	// flags
	active := flag.String("active", "", "active profile")
	flag.Parse()

	// info
	pwd, _ := os.Getwd()
	log.Info("PWD:", pwd)

	activeProfile := ".env"
	if *active != "" {
		activeProfile = fmt.Sprintf(".env-%s", *active)
	}
	log.Info("Active Profile:", activeProfile)

	// 加载配置文件
	config.LoadEnvFile(activeProfile)
}
