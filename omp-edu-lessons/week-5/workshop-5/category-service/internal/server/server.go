package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"

	api "github.com/ozonmp/week-5-workshop/category-service/internal/app/category-service"
	"github.com/ozonmp/week-5-workshop/category-service/internal/config"
	"github.com/ozonmp/week-5-workshop/category-service/internal/pkg/logger"
	mwserver "github.com/ozonmp/week-5-workshop/category-service/internal/pkg/mw/server"
	"github.com/ozonmp/week-5-workshop/category-service/internal/service/category"
	"github.com/ozonmp/week-5-workshop/category-service/internal/service/task"
	desc "github.com/ozonmp/week-5-workshop/category-service/pkg/category-service"
)

type GrpcServer struct {
	categoryService category.Service
	taskService     task.Service
}

func NewGrpcServer(
	categoryService category.Service,
	taskService task.Service,
) *GrpcServer {
	return &GrpcServer{
		categoryService: categoryService,
		taskService:     taskService,
	}
}

func (s *GrpcServer) Start(cfg *config.Config) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	gatewayAddr := fmt.Sprintf("%s:%v", cfg.Gateway.Host, cfg.Gateway.Port)
	swaggerAddr := fmt.Sprintf("%s:%v", cfg.Swagger.Host, cfg.Swagger.Port)
	grpcAddr := fmt.Sprintf("%s:%v", cfg.Grpc.Host, cfg.Grpc.Port)

	gatewayServer := createGatewayServer(grpcAddr, gatewayAddr, cfg.Gateway.AllowedCORSOrigins)
	swaggerServer, err := createSwaggerServer(gatewayAddr, swaggerAddr, cfg.Swagger.Filepath)
	if err != nil {
		return err
	}

	go func() {
		logger.InfoKV(ctx, fmt.Sprintf("Gateway server is running on %s", gatewayAddr))
		if err := gatewayServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.ErrorKV(ctx, "Failed running gateway server")
			cancel()
		}
	}()

	go func() {
		logger.InfoKV(ctx, fmt.Sprintf("Swagger server is running on %s", swaggerAddr))
		if err := swaggerServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.ErrorKV(ctx, "Failed running swagger server")
			cancel()
		}
	}()

	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}
	defer l.Close()

	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: time.Duration(cfg.Grpc.MaxConnectionIdle) * time.Minute,
			Timeout:           time.Duration(cfg.Grpc.Timeout) * time.Second,
			MaxConnectionAge:  time.Duration(cfg.Grpc.MaxConnectionAge) * time.Minute,
			Time:              time.Duration(cfg.Grpc.Timeout) * time.Minute,
		}),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_opentracing.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_recovery.UnaryServerInterceptor(),
			mwserver.GRPCUnauthenticatedRequest,
		)),
	)

	desc.RegisterCategoryServiceServer(grpcServer, api.NewCategoryService(
		s.categoryService,
		s.taskService,
	))

	go func() {
		logger.InfoKV(ctx, fmt.Sprintf("GRPC Server is listening on: %s", grpcAddr))
		if err := grpcServer.Serve(l); err != nil {
			logger.FatalKV(ctx, "Failed running gRPC server", "err", err)
		}
	}()

	if cfg.Project.Debug {
		reflection.Register(grpcServer)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		logger.InfoKV(ctx, fmt.Sprintf("signal.Notify: %v", v))
	case done := <-ctx.Done():
		logger.InfoKV(ctx, fmt.Sprintf("ctx.Done: %v", done))
	}

	if err := gatewayServer.Shutdown(ctx); err != nil {
		logger.ErrorKV(ctx, "gatewayServer.Shutdown")
	} else {
		logger.InfoKV(ctx, "gatewayServer shut down correctly")
	}

	grpcServer.GracefulStop()
	logger.InfoKV(ctx, fmt.Sprintf("grpcServer shut down correctly"))

	return nil
}
