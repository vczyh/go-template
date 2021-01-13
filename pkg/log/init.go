package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var logger *zap.Logger
var sugar *zap.SugaredLogger

func ConfigLog() {
	// writer
	ws := getWriter()
	// encoder
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	core := zapcore.NewCore(encoder, ws, zapcore.DebugLevel)

	logger = zap.New(core)
	sugar = logger.Sugar()
}

func getWriter() zapcore.WriteSyncer {
	l := &lumberjack.Logger{
		Filename:   "./test.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		LocalTime:  true,
		Compress:   false,
	}
	//l.Rotate()
	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(l), zapcore.AddSync(os.Stdout))
}

// DEBUG
func Debug(args ...interface{}) {
	sugar.Debug(args)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	sugar.Debugw(msg, keysAndValues...)
}

// INFO
func Info(args ...interface{}) {
	sugar.Info(args)
}

func Infow(msg string, keysAndValues ...interface{}) {
	sugar.Infow(msg, keysAndValues...)
}

// WARN
func Warn(args ...interface{}) {
	sugar.Warn(args)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	sugar.Warnw(msg, keysAndValues...)
}

// ERROR
func Error(args ...interface{}) {
	sugar.Error(args)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	sugar.Errorw(msg, keysAndValues...)
}

// Fatal
func Fatal(args ...interface{}) {
	sugar.Fatal(args)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	sugar.Fatalw(msg, keysAndValues...)
}