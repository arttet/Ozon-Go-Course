package main

import (
	"database/sql"
	"log"

	"github.com/go-testfixtures/testfixtures/v3"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	db, err := sql.Open("pgx", "postgres://user:password@localhost:5432/db")
	if err != nil {
		log.Fatalf("sql.Open(): %v", err)
	}
	defer db.Close()

	fixtures, err := testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("postgres"),
		testfixtures.Directory(
			"7-postgresql-part-1/4-fixtures/fixtures",
		),
		testfixtures.DangerousSkipTestDatabaseCheck(), // allow non-test DB name
	)
	if err != nil {
		log.Fatalf("testfixtures.New(): %v", err)
	}

	err = fixtures.Load()
	if err != nil {
		log.Fatalf("fixtures.Load(): %v", err)
	}
}
