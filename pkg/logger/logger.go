package logger

import (
	"context"
	"go.uber.org/zap"
)

const LoggerKey = "logger"
const requestID = "requestID"
const ServiceName = "service"

type Logger interface {
	Info(ctx context.Context, msg string, field ...zap.Field)
	Error(ctx context.Context, msg string, field ...zap.Field)
}

type logger struct {
	serviceName string
	logger      *zap.Logger
}

func (l logger) Info(ctx context.Context, msg string, field ...zap.Field) {
	field = append(field, zap.String(ServiceName, l.serviceName))
	if ctx.Value(requestID) != nil {
		field = append(field, zap.String(requestID, ctx.Value(requestID).(string)))
	}
	l.logger.Info(msg, field...)
}

func (l logger) Error(ctx context.Context, msg string, field ...zap.Field) {
	field = append(field, zap.String(ServiceName, l.serviceName))
	if ctx.Value(requestID) != nil {
		field = append(field, zap.String(requestID, ctx.Value(requestID).(string)))
	}
	l.logger.Info(msg, field...)
}

func New(serviceName string) Logger {
	zapLogger, _ := zap.NewProduction()
	defer zapLogger.Sync()
	return &logger{
		serviceName: serviceName,
		logger:      zapLogger,
	}

}

func GetLoggerFromCtx(ctx context.Context) Logger {
	return ctx.Value(LoggerKey).(Logger)

}
