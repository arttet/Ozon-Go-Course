package main

import (
	"context"
	"log"

	"github.com/ozonmp/week-4-workshop/category-service/internal/config"
	"github.com/ozonmp/week-4-workshop/category-service/internal/service/database"
	"github.com/ozonmp/week-4-workshop/category-service/migrations"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pressly/goose/v3"
)

func main() {
	err := config.ReadConfigYML("config.yml")
	if err != nil {
		log.Fatalf("config.ReadConfigYML(): %v", err)
	}
	cfg := config.GetConfigInstance()

	ctx := context.Background()
	db, err := database.New(ctx, cfg.Database.DSN)
	if err != nil {
		log.Fatalf("database.New(): %v", err)
	}

	goose.SetBaseFS(migrations.EmbedFS)

	err = goose.Up(db.DB, ".")
	if err != nil {
		log.Fatalf("goose.Up(): %v", err)
	}
}
