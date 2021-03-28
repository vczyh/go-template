package demo

import (
	"go-template/pkg/config"
	"go-template/pkg/log"
)

// log
var l *log.Logger

func WithLogger(logger *log.Logger) {
	l = logger
}

// config
var c *config.Config

func WithConfig(config *config.Config) {
	c = config
}
