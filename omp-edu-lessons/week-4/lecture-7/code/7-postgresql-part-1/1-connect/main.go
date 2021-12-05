package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	fmt.Println()

	// "user=user password=password host=localhost port=5432 database=db sslmode=disable"
	db, err := sql.Open("pgx", "postgres://user:password@localhost:5432/db")
	if err != nil {
		log.Fatalf("sql.Open(): %v", err)
	}

	_, _ = fmt.Scanln()

	// Только сейчас создается подключение
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatalf("db.PingContext(): %v", err)
	}

	fmt.Println("Done")
}
