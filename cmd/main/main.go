package main

import (
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"os/signal"
	"syscall"
	"yandexlms/internal/config"
	"yandexlms/internal/transport/grpc"
	"yandexlms/pkg/db/cache"
	"yandexlms/pkg/logger"
)

const (
	serviceName = "lyceum"
)

func main() {
	ctx := context.Background()

	mainLogger := logger.New(serviceName)
	ctx = context.WithValue(ctx, logger.LoggerKey, mainLogger)

	cfg := config.New()
	if cfg == nil {
		panic("Fatal lo load config")
	}

	//dbCfg := postgres.Config{"postgres", "123", "localhost", "5432", "yandex"}
	//db, err := postgres.New(dbCfg)
	//if err != nil {
	//	println(err)
	//}
	//fmt.Println(db)

	redis := cache.New(cfg.RedisConfig)
	fmt.Println(redis.Ping(ctx))

	graceCh := make(chan os.Signal, 1)
	signal.Notify(graceCh, syscall.SIGINT, syscall.SIGTERM)

	grpcServer, err := grpc.New(ctx, cfg.GRPCServerPort, cfg.RestServerPort)
	if err != nil {
		mainLogger.Error(ctx, err.Error())
		return
	}

	go func() {
		err = grpcServer.Start(ctx)
		if err != nil {
			mainLogger.Error(ctx, err.Error())
		}
	}()

	<-graceCh
	err = grpcServer.Stop(ctx)
	if err != nil {
		mainLogger.Error(ctx, err.Error())
	}
	mainLogger.Info(ctx, "Server stopped")
}
