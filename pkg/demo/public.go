package demo

import "go-template/pkg/log"

var logger *log.Logger

func WithLogger(l *log.Logger) {
	logger = l
}
