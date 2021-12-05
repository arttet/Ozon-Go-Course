package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"

	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/opentracing/opentracing-go"
	"github.com/ozonmp/week-5-workshop/product-service/internal/pkg/db"
	"github.com/ozonmp/week-5-workshop/product-service/internal/pkg/logger"
	gelf "github.com/snovichkov/zap-gelf"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"

	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"

	grpc_category_service "github.com/ozonmp/week-5-workshop/category-service/pkg/category-service"

	"github.com/ozonmp/week-5-workshop/product-service/internal/config"
	mwclient "github.com/ozonmp/week-5-workshop/product-service/internal/pkg/mw/client"
	"github.com/ozonmp/week-5-workshop/product-service/internal/server"
	product_service "github.com/ozonmp/week-5-workshop/product-service/internal/service/product"
)

func main() {
	ctx := context.Background()

	if err := config.ReadConfigYML("config.yml"); err != nil {
		logger.FatalKV(ctx, "Failed init configuration", "err", err)
	}
	cfg := config.GetConfigInstance()

	flag.Parse()

	syncLogger := initLogger(ctx, cfg)
	defer syncLogger()

	closer := initTracer(ctx, cfg)
	defer closer.Close()

	logger.InfoKV(ctx, fmt.Sprintf("Starting service: %s", cfg.Project.Name),
		"version", cfg.Project.Version,
		"commitHash", cfg.Project.CommitHash,
		"debug", cfg.Project.Debug,
		"environment", cfg.Project.Environment,
	)

	categoryServiceConn, err := grpc.DialContext(
		context.Background(),
		cfg.CategoryServiceAddr,
		grpc.WithInsecure(),
		grpc.WithChainUnaryInterceptor(
			mwclient.AddAppInfoUnary,
			grpc_opentracing.UnaryClientInterceptor(),
		),
	)
	if err != nil {
		logger.ErrorKV(ctx, "failed to create client", "err", err)
	}

	conn, err := db.ConnectDB(&cfg.DB)
	if err != nil {
		logger.FatalKV(ctx, "sql.Open() error", "err", err)
	}
	defer conn.Close()

	categoryServiceClient := grpc_category_service.NewCategoryServiceClient(categoryServiceConn)

	productService := product_service.NewService(categoryServiceClient, conn)

	if err := server.NewGrpcServer(productService).Start(&cfg); err != nil {
		logger.ErrorKV(ctx, "Failed creating gRPC server", "err", err)

		return
	}
}

func initLogger(ctx context.Context, cfg config.Config) (syncFn func()) {
	loggingLevel := zap.InfoLevel
	if cfg.Project.Debug {
		loggingLevel = zap.DebugLevel
	}

	consoleCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		os.Stderr,
		zap.NewAtomicLevelAt(loggingLevel),
	)

	gelfCore, err := gelf.NewCore(
		gelf.Addr(cfg.Telemetry.GraylogPath),
		gelf.Level(loggingLevel),
	)
	if err != nil {
		logger.FatalKV(ctx, "sql.Open() error", "err", err)
	}

	notSugaredLogger := zap.New(zapcore.NewTee(consoleCore, gelfCore))

	sugaredLogger := notSugaredLogger.Sugar()
	logger.SetLogger(sugaredLogger.With(
		"service", cfg.Project.ServiceName,
	))

	return func() {
		notSugaredLogger.Sync()
	}
}

func initTracer(ctx context.Context, cfg config.Config) (closer io.Closer) {
	// Sample configuration for testing. Use constant sampling to sample every trace
	// and enable LogSpan to log every span via configured Logger.
	jcfg := jaegercfg.Configuration{
		ServiceName: cfg.Project.ServiceName,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}

	// Initialize tracer with a logger and a metrics factory
	tracer, closer, err := jcfg.NewTracer()
	if err != nil {
		logger.FatalKV(ctx, "cfg.NewTracer()", "err", err)
	}
	// Set the singleton opentracing.Tracer with the Jaeger tracer.
	opentracing.SetGlobalTracer(tracer)
	return closer
}
