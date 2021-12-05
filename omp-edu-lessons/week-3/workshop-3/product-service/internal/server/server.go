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

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"

	api "github.com/ozonmp/week-3-workshop/product-service/internal/app/product-service"
	"github.com/ozonmp/week-3-workshop/product-service/internal/config"
	product_service "github.com/ozonmp/week-3-workshop/product-service/internal/service/product"
	desc "github.com/ozonmp/week-3-workshop/product-service/pkg/product-service"
)

type GrpcServer struct {
	productService *product_service.Service
}

func NewGrpcServer(productService *product_service.Service) *GrpcServer {
	return &GrpcServer{
		productService: productService,
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
		log.Info().Msgf("Gateway server is running on %s", gatewayAddr)
		if err := gatewayServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error().Err(err).Msg("Failed running gateway server")
			cancel()
		}
	}()

	go func() {
		log.Info().Msgf("Swagger server is running on %s", swaggerAddr)
		if err := swaggerServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error().Err(err).Msg("Failed running swagger server")
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
			grpcrecovery.UnaryServerInterceptor(),
		)),
	)

	desc.RegisterProductServiceServer(grpcServer, api.NewProductService(s.productService))

	go func() {
		log.Info().Msgf("GRPC Server is listening on: %s", grpcAddr)
		if err := grpcServer.Serve(l); err != nil {
			log.Fatal().Err(err).Msg("Failed running gRPC server")
		}
	}()

	if cfg.Project.Debug {
		reflection.Register(grpcServer)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		log.Info().Msgf("signal.Notify: %v", v)
	case done := <-ctx.Done():
		log.Info().Msgf("ctx.Done: %v", done)
	}

	if err := gatewayServer.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("gatewayServer.Shutdown")
	} else {
		log.Info().Msg("gatewayServer shut down correctly")
	}

	grpcServer.GracefulStop()
	log.Info().Msgf("grpcServer shut down correctly")

	return nil
}
