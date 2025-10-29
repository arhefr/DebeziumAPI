package v1

import (
	"context"
	"debez/pkg/logger"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func RegistLoggerInterceptor(log logger.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		ctx = logger.WithLogger(ctx, log)
		return handler(ctx, req)
	}
}

func LoggerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		log := logger.FromContext(ctx)
		start := time.Now()
		resp, err = handler(ctx, req)
		code := status.Code(err)

		fields := []zap.Field{
			zap.String("Method", info.FullMethod),
			zap.String("StatusCode", code.String()),
			zap.String("Time", time.Since(start).String()),
		}

		switch code {
		case codes.OK:
			log.Info(ctx, "gRPC request succecs", fields...)
		default:
			log.Info(ctx, "gRPC request fail", fields...)
		}
		return resp, err
	}
}
