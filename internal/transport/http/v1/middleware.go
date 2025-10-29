package v1

import (
	"context"
	"debez/pkg/logger"
	"net/http"

	"github.com/google/uuid"
)

const (
	HeaderRequestIDKey string = "X-Request-ID"
	HeaderTraceIDKey   string = "X-Trace-ID"

	CtxRequestIDKey string = "x-request-id"
	CtxTraceIDKey   string = "x-trace-id"
)

func AddMetadata(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		traceID := r.Header.Get(string(HeaderTraceIDKey))
		if traceID == "" {
			traceID = uuid.NewString()
		}

		requestID := uuid.NewString()

		r.Header.Set(HeaderTraceIDKey, traceID)
		r.Header.Set(HeaderRequestIDKey, requestID)

		ctx := r.Context()
		ctx = context.WithValue(ctx, CtxTraceIDKey, traceID)
		ctx = context.WithValue(ctx, CtxRequestIDKey, requestID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RegistLoggerMiddleware(log logger.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := logger.WithLogger(r.Context(), log)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log := logger.FromContext(ctx)
		log.Info(ctx, "HTTP request succecs")
		next.ServeHTTP(w, r)

	})
}
