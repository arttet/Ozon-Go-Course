package main

import (
	"log"

	"github.com/ozonmp/week-5-workshop/product-service/internal/config"
	"github.com/ozonmp/week-5-workshop/product-service/internal/pkg/db"
	"github.com/ozonmp/week-5-workshop/product-service/migrations"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pressly/goose/v3"
)

func main() {
	err := config.ReadConfigYML("config.yml")
	if err != nil {
		log.Fatalf("config.ReadConfigYML(): %v", err)
	}
	cfg := config.GetConfigInstance()

	conn, err := db.ConnectDB(&cfg.DB)
	if err != nil {
		log.Fatalf("sql.Open() error: %v", err)
	}
	defer conn.Close()

	goose.SetBaseFS(migrations.EmbedFS)

	err = goose.Up(conn.DB, ".")
	if err != nil {
		log.Fatalf("goose.Up(): %v", err)
	}
}
