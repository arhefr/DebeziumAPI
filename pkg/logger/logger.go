package logger

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	loggerKey string = "logger"

	HeaderRequestIDKey string = "X-Request-ID"
	HeaderTraceIDKey   string = "X-Trace-ID"

	CtxRequestIDKey string = "x-request-id"
	CtxTraceIDKey   string = "x-trace-id"
)

type Logger interface {
	Warn(ctx context.Context, msg string, fields ...zap.Field)
	Debug(ctx context.Context, msg string, fields ...zap.Field)
	Error(ctx context.Context, msg string, fields ...zap.Field)
	Fatal(ctx context.Context, msg string, fields ...zap.Field)
	Info(ctx context.Context, msg string, fields ...zap.Field)
}

func NewLogger(logLevel zapcore.Level) (Logger, error) {
	loggerCfg := zap.NewProductionConfig()

	loggerCfg.Level = zap.NewAtomicLevelAt(logLevel)

	logger, err := loggerCfg.Build()
	if err != nil {
		return nil, err
	}
	defer logger.Sync()

	return &L{z: logger}, nil
}

type L struct {
	z *zap.Logger
}

func (l L) Debug(ctx context.Context, msg string, fields ...zap.Field) {

	requestID, traceID := ctx.Value(CtxRequestIDKey), ctx.Value(CtxTraceIDKey)

	fields = append(fields, zap.Any(CtxRequestIDKey, requestID), zap.Any(CtxTraceIDKey, traceID))

	l.z.Debug(msg, fields...)
}

func (l L) Info(ctx context.Context, msg string, fields ...zap.Field) {

	requestID, traceID := ctx.Value(CtxRequestIDKey), ctx.Value(CtxTraceIDKey)

	fields = append(fields, zap.Any(CtxRequestIDKey, requestID), zap.Any(CtxTraceIDKey, traceID))

	l.z.Info(msg, fields...)
}

func (l L) Warn(ctx context.Context, msg string, fields ...zap.Field) {
	requestID, traceID := ctx.Value(CtxRequestIDKey), ctx.Value(CtxTraceIDKey)

	fields = append(fields, zap.Any(CtxRequestIDKey, requestID), zap.Any(CtxTraceIDKey, traceID))

	l.z.Warn(msg, fields...)
}

func (l L) Error(ctx context.Context, msg string, fields ...zap.Field) {
	requestID, traceID := ctx.Value(CtxRequestIDKey), ctx.Value(CtxTraceIDKey)

	fields = append(fields, zap.Any(CtxRequestIDKey, requestID), zap.Any(CtxTraceIDKey, traceID))

	l.z.Error(msg, fields...)
}

func (l L) Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	requestID, traceID := ctx.Value(CtxRequestIDKey), ctx.Value(CtxTraceIDKey)

	fields = append(fields, zap.Any(CtxRequestIDKey, requestID), zap.Any(CtxTraceIDKey, traceID))

	l.z.Fatal(msg, fields...)
}

func WithLogger(ctx context.Context, l Logger) context.Context {
	return context.WithValue(ctx, loggerKey, l)
}

func FromContext(ctx context.Context) Logger {
	if log, ok := ctx.Value(loggerKey).(*L); ok {
		return log
	}

	return nil
}
