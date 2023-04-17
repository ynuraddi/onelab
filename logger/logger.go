package logger

import (
	"context"
	"log"

	"go.uber.org/zap"
)

type loggerType int

const (
	requestIDKey loggerType = iota
	sessionIDKey
)

var logger zap.SugaredLogger

func init() {
	c := zap.NewProductionConfig()
	c.OutputPaths = []string{"./tmp/logs.log"}
	l, err := c.Build()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()

	logger = *l.Sugar()
}

func WithRqId(ctx context.Context, reqID string) context.Context {
	return context.WithValue(ctx, requestIDKey, reqID)
}

func WithSessionId(ctx context.Context, sessionId string) context.Context {
	return context.WithValue(ctx, sessionIDKey, sessionId)
}

func Info(ctx context.Context, msg string) {
	l := Logger(ctx)
	l.With(msg).Info()
}

func Logger(ctx context.Context) zap.SugaredLogger {
	newLogger := &logger
	if ctx != nil {
		if ctxRqId, ok := ctx.Value(requestIDKey).(string); ok {
			newLogger = newLogger.With(zap.String("rqId", ctxRqId))
		}
		if ctxSessionId, ok := ctx.Value(sessionIDKey).(string); ok {
			newLogger = newLogger.With(zap.String("sessionId", ctxSessionId))
		}
	}

	return *newLogger
}
