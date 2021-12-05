package main

import (
	"context"
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

func main() {
	fmt.Println()

	db, openErr := sqlx.Open("pgx", "postgres://user:password@localhost:5432/db")
	if openErr != nil {
		log.Fatalf("sqlx.Open(): %v", openErr)
	}

	sb := psql.Select("id", "name").
		From("products")

	categoryIDs := []int64{1, 3, 5}
	if len(categoryIDs) != 0 {
		sb = sb.Join("categories c on c.id = p.category_id")
		sb = sb.Where(sq.Eq{"c.id": categoryIDs})
	}

	and := sq.And{
		sq.Eq{"name": "Audi"},   // name = 'Audi'
		sq.Like{"name": "%cap"}, // name LIKE '%cap'
	}
	and = append(and, sq.NotEq{"created_at": nil}) // created_at IS NOTNULL

	sb = sb.Where(and)

	query, args, err := sb.ToSql()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	type product struct {
		ID   int64  `db:"id"`
		Name string `db:"name"`
	}

	var products []product
	err = db.SelectContext(ctx, &products, query, args...)
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range products {
		fmt.Printf("%d — %s\n", p.ID, p.Name)
	}
}
