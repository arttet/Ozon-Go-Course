package main

import (
	"database/sql"
	"embed"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	db, err := sql.Open("pgx", "postgres://user:password@localhost:5432/db")
	if err != nil {
		log.Fatalf("sql.Open(): %v", err)
	}
	defer db.Close()

	goose.SetBaseFS(embedMigrations)

	const cmd = "up"
	err = goose.Run(cmd, db, "migrations")
	if err != nil {
		log.Fatalf("goose.Status(): %v", err)
	}
}
