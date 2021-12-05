package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	_ "github.com/jackc/pgx/v4/stdlib"
	gelf "github.com/snovichkov/zap-gelf"

	jaegercfg "github.com/uber/jaeger-client-go/config"

	"github.com/ozonmp/week-5-workshop/category-service/internal/config"
	"github.com/ozonmp/week-5-workshop/category-service/internal/pkg/logger"
	"github.com/ozonmp/week-5-workshop/category-service/internal/server"
	"github.com/ozonmp/week-5-workshop/category-service/internal/service/category"
	cat_repository "github.com/ozonmp/week-5-workshop/category-service/internal/service/category/repository"
	"github.com/ozonmp/week-5-workshop/category-service/internal/service/database"
	"github.com/ozonmp/week-5-workshop/category-service/internal/service/task"
	task_repository "github.com/ozonmp/week-5-workshop/category-service/internal/service/task/repository"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

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

	initCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := database.New(initCtx, cfg.Database.DSN)
	if err != nil {
		logger.ErrorKV(ctx, "failed to create client", "err", err)
	}

	categoryRepository := cat_repository.New(db)
	categoryService := category.New(categoryRepository)

	taskRepository := task_repository.New(db)
	taskService := task.New(taskRepository, db)

	if err := server.NewGrpcServer(
		categoryService,
		taskService,
	).Start(&cfg); err != nil {
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
