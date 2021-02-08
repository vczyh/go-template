package demo

import "go-template/pkg/log"

// log
var logger *log.Logger

func WithLogger(l *log.Logger) {
	logger = l
}
