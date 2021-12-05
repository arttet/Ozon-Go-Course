package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
)

func main() {
	fmt.Println()

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://user:password@localhost:5432/db")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer conn.Close(ctx)

	var pgVersion string
	err = conn.QueryRow(ctx, "select version()").Scan(&pgVersion)
	if err != nil {
		log.Fatalf("QueryRow failed: %v", err)
	}

	fmt.Println("Postgres version:", pgVersion)
}
