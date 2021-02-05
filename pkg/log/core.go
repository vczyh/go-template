package log

import (
	"context"
	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
)

type Logger struct {
	serviceName string
	contextKeys []string

	base  *zap.Logger
	sugar *zap.SugaredLogger
}

func NewLogger(serviceName string, contextKeys []string, writers ... io.Writer) *Logger {
	io.MultiWriter(writers...)
	return &Logger{
		serviceName: serviceName,
		contextKeys: contextKeys,
	}
}

func (l *Logger) Debug(ctx context.Context, msg string, keysAndValues ...interface{}) {
	l.sugar.Debugw(msg, contextKeysAndValues(ctx, keysAndValues...)...)
}

func contextKeysAndValues(ctx context.Context, keysAndValues ...interface{}) []interface{} {
	var kvs []interface{}
	for _, key := range contextKeys {
		kvs = append(kvs, key, ctx.Value(key))
	}
	return append(kvs, keysAndValues...)
}

//
func NewFileWriter(file string, maxSize int, maxBackups int, maxAge int) io.WriteCloser {
	return &lumberjack.Logger{
		Filename:   file,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		LocalTime:  true,
		Compress:   false,
	}
}
