package main

import (
	"context"
	"flag"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/ozonmp/week-4-workshop/category-service/internal/config"
	"github.com/ozonmp/week-4-workshop/category-service/internal/server"
	"github.com/ozonmp/week-4-workshop/category-service/internal/service/category"
	cat_repository "github.com/ozonmp/week-4-workshop/category-service/internal/service/category/repository"
	"github.com/ozonmp/week-4-workshop/category-service/internal/service/database"
	"github.com/ozonmp/week-4-workshop/category-service/internal/service/task"
	task_repository "github.com/ozonmp/week-4-workshop/category-service/internal/service/task/repository"
)

func main() {
	if err := config.ReadConfigYML("config.yml"); err != nil {
		log.Fatal().Err(err).Msg("Failed init configuration")
	}
	cfg := config.GetConfigInstance()

	flag.Parse()

	log.Info().
		Str("version", cfg.Project.Version).
		Str("commitHash", cfg.Project.CommitHash).
		Bool("debug", cfg.Project.Debug).
		Str("environment", cfg.Project.Environment).
		Msgf("Starting service: %s", cfg.Project.Name)

	// default: zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if cfg.Project.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	initCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := database.New(initCtx, cfg.Database.DSN)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed init configuration")
	}

	categoryRepository := cat_repository.New(db)
	categoryService := category.New(categoryRepository)

	taskRepository := task_repository.New(db)
	taskService := task.New(taskRepository, db)

	if err := server.NewGrpcServer(
		categoryService,
		taskService,
	).Start(&cfg); err != nil {
		log.Error().Err(err).Msg("Failed creating gRPC server")

		return
	}
}
