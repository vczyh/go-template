package log

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
)

type Logger struct {
	serviceName string
	contextKeys []string

	base  *zap.Logger
	sugar *zap.SugaredLogger
}

func NewLogger(serviceName string, contextKeys []string, writers ...io.Writer) *Logger {
	logger := &Logger{
		serviceName: serviceName,
		contextKeys: contextKeys,
	}
	// encoder
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	level := zap.NewAtomicLevelAt(zap.DebugLevel)
	ws := io.MultiWriter(writers...)

	core := zapcore.NewCore(
		//zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(ws)),
		level,
	)
	l := zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	)
	l = l.Named(logger.serviceName)

	logger.base = l
	logger.sugar = l.Sugar()

	return logger
}

func (l *Logger) name() {

}

func (l *Logger) Debug(ctx context.Context, msg string, keysAndValues ...interface{}) {
	l.sugar.Debugw(msg, l.contextKeysAndValues(ctx, keysAndValues...)...)
}

func (l *Logger) contextKeysAndValues(ctx context.Context, keysAndValues ...interface{}) []interface{} {
	if ctx == nil {
		return keysAndValues
	}
	var kvs []interface{}
	for _, key := range l.contextKeys {
		kvs = append(kvs, key, ctx.Value(key))
	}
	return append(kvs, keysAndValues...)
}

func NewRotate(file string, maxSize int, maxBackups int, maxAge int) io.WriteCloser {
	return &lumberjack.Logger{
		Filename:   file,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		LocalTime:  true,
		Compress:   false,
	}
}
