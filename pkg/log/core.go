package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
)

type Logger struct {
	name string

	base  *zap.Logger
	sugar *zap.SugaredLogger
}

func NewLogger(name string, level string, writers ...io.Writer) *Logger {
	logger := &Logger{
		name: name,
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

	var lev zap.AtomicLevel
	switch level {
	case "DEBUG":
		lev = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "INFO":
		lev = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "WARN":
		lev = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "ERROR":
		lev = zap.NewAtomicLevelAt(zap.ErrorLevel)
	default:
		lev = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	ws := io.MultiWriter(writers...)

	core := zapcore.NewCore(
		//zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(ws)),
		lev,
	)
	l := zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	)
	l = l.Named(logger.name)

	logger.base = l
	logger.sugar = l.Sugar()

	return logger
}

func (l *Logger) Debug(msg string, kvs ...interface{}) {
	l.sugar.Debugw(msg, kvs...)
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
