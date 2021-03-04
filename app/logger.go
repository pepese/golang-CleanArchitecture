package app

import (
	"context"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type key int

const loggerKey = key(1)

var (
	desugaredLogger     *zap.Logger
	desugaredLoggerOnce sync.Once
	logger              *zap.SugaredLogger
	loggerOnce          sync.Once
)

func getDesugaredLogger() *zap.Logger {
	desugaredLoggerOnce.Do(func() {
		config := zap.Config{
			Level:    zap.NewAtomicLevelAt(zapcore.InfoLevel),
			Encoding: "json",
			EncoderConfig: zapcore.EncoderConfig{
				MessageKey:     "message",
				LevelKey:       "level",
				TimeKey:        "timestamp",
				NameKey:        "name",
				CallerKey:      "caller",
				StacktraceKey:  "stacktrace",
				EncodeLevel:    zapcore.LowercaseLevelEncoder,
				EncodeTime:     zapcore.ISO8601TimeEncoder,
				EncodeDuration: zapcore.StringDurationEncoder,
				EncodeCaller:   zapcore.ShortCallerEncoder,
			},
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stderr"},
		}
		zapLogger, err := config.Build()
		if err != nil {
			panic(err)
		}
		desugaredLogger = zapLogger.With(
			zap.String("app", Config().AppName),
			zap.String("version", Config().AppVersion),
		)
	})
	return desugaredLogger
}

func Logger() *zap.SugaredLogger {
	tmpLogger := getDesugaredLogger()
	loggerOnce.Do(func() {
		logger = tmpLogger.Sugar()
	})
	return logger
}

func LoggerFromContext(c context.Context) *zap.SugaredLogger {
	if c == nil {
		return Logger()
	}
	if zapLogger, ok := c.Value(loggerKey).(*zap.SugaredLogger); ok {
		return zapLogger
	}
	return Logger()

}

func SetLoggerToContext(c context.Context, logger *zap.SugaredLogger) context.Context {
	return context.WithValue(c, loggerKey, logger)
}

func LoggerWithKeyValue(key string, value string) *zap.SugaredLogger {
	tmpLogger := getDesugaredLogger()
	tmpLogger = tmpLogger.With(
		zap.String(key, value),
	)
	return tmpLogger.Sugar()
}
