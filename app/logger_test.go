package app

import (
	"context"
	"testing"
)

type testStruct struct {
	param1 string
	param2 string
}

func TestLogger(t *testing.T) {
	ctx := context.Background()
	t.Run("Logging message", func(t *testing.T) {
		logger := Logger()
		logger.Info("hoge")
	})
	t.Run("Logging message with reqId", func(t *testing.T) {
		logger := LoggerWithKeyValue("reqId", "aaaaaaaaaa")
		logger.Info("fuge")
	})
	t.Run("Set/Get logger to context", func(t *testing.T) {
		ctx = SetLoggerToContext(ctx, LoggerWithKeyValue("reqId", "bbbbbbbbbb"))
		logger := LoggerFromContext(ctx)
		logger.Info("guhe")
	})
	t.Run("Logging struxt message", func(t *testing.T) {
		logger := Logger()
		msg := testStruct{param1: "hoge", param2: "huge"}
		logger.Info(msg)
	})
}
