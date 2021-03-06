package demo

import (
	"blog-y/pkg/common/config"
	"blog-y/pkg/common/log"
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
