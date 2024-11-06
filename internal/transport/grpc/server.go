package grpc

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"strconv"
	client "yandexlms/pkg/api/order"
	"yandexlms/pkg/logger"
)

type Server struct {
	grpcServer *grpc.Server
	restServer *http.Server
	listener   net.Listener
}

func New(ctx context.Context, port, restPort int) (*Server, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(LoggerInterceptor),
	}

	grpcServer := grpc.NewServer(opts...)
	client.RegisterOrderServiceServer(grpcServer, NewOrderService())

	restSrv := runtime.NewServeMux()
	if err := client.RegisterOrderServiceHandlerServer(context.Background(), restSrv, NewOrderService()); err != nil {
		return nil, err
	}
	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", restPort),
		Handler: restSrv,
	}

	return &Server{grpcServer: grpcServer, restServer: httpServer, listener: lis}, nil
}

func (s *Server) Start(ctx context.Context) error {
	eg := errgroup.Group{}

	eg.Go(func() error {
		logger.GetLoggerFromCtx(ctx).Info(ctx, "grpc server started", zap.String("port", strconv.Itoa(s.listener.Addr().(*net.TCPAddr).Port)))
		return s.grpcServer.Serve(s.listener)
	})

	eg.Go(func() error {
		logger.GetLoggerFromCtx(ctx).Info(ctx, "rest server started", zap.String("port", s.restServer.Addr))
		return s.restServer.ListenAndServe()
	})

	return eg.Wait()
}

func (s *Server) Stop(ctx context.Context) error {
	s.grpcServer.Stop()
	l := logger.GetLoggerFromCtx(ctx)
	l.Info(ctx, "gRPC server stopped")
	err := s.restServer.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("failed to stop rest server: %w", err)
	}
	l.Info(ctx, "Rest server stopped")
	return nil
}
