package log

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

var (
	logger *zap.Logger
	sugar  *zap.SugaredLogger
)

var contextKeys []string

func ConfigLog(keys []string) {
	// writer
	w := getWriter()
	ws := zapcore.NewMultiWriteSyncer(zapcore.AddSync(w))
	// encoder
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	core := zapcore.NewCore(encoder, ws, zapcore.DebugLevel)

	logger = zap.New(core)
	sugar = logger.Sugar()

	// gin
	gin.DefaultWriter = w

	contextKeys = keys
}

func getWriter() io.Writer {
	logPath := os.Getenv("LOG_PATH")
	l := &lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		LocalTime:  true,
		Compress:   false,
	}
	return io.MultiWriter(l, os.Stdout)
}

func Debugw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	sugar.Debugw(msg, contextInfo(ctx, keysAndValues...)...)
}

func Infow(ctx context.Context, msg string, keysAndValues ...interface{}) {
	sugar.Infow(msg, contextInfo(ctx, keysAndValues)...)
}

func contextInfo(ctx context.Context, keysAndValues ...interface{}) []interface{} {
	var kvs []interface{}
	for _, key := range contextKeys {
		kvs = append(kvs, key, ctx.Value(key))
	}
	return append(kvs, keysAndValues...)
}

// DEBUG
func debug(args ...interface{}) {
	sugar.Debug(args)
}

func debugw(msg string, keysAndValues ...interface{}) {
	sugar.Debugw(msg, keysAndValues...)
}

// INFO
func info(args ...interface{}) {
	sugar.Info(args)
}

func infow(msg string, keysAndValues ...interface{}) {
	sugar.Infow(msg, keysAndValues...)
}

// WARN
func warn(args ...interface{}) {
	sugar.Warn(args)
}

func warnw(msg string, keysAndValues ...interface{}) {
	sugar.Warnw(msg, keysAndValues...)
}

// ERROR
func error(args ...interface{}) {
	sugar.Error(args)
}

func errorw(msg string, keysAndValues ...interface{}) {
	sugar.Errorw(msg, keysAndValues...)
}

// FATAL
func fatal(args ...interface{}) {
	sugar.Fatal(args)
}

func fatalw(msg string, keysAndValues ...interface{}) {
	sugar.Fatalw(msg, keysAndValues...)
}
