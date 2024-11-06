package grpc

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"yandexlms/pkg/logger"
)

func LoggerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	logs := logger.GetLoggerFromCtx(ctx)
	logs.Info(ctx, "request started", zap.String("method", info.FullMethod))
	return handler(ctx, req)
}
