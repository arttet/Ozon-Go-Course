package main

import (
	"context"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type product struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}

func main() {
	db, openErr := sqlx.Open("pgx", "postgres://user:password@localhost:5432/db")
	if openErr != nil {
		log.Fatalf("sqlx.Open(): %v", openErr)
	}

	ctx := context.Background()

	fmt.Println()
	fmt.Println("GetContext")

	{
		query := `SELECT p.id, p.name, p.created_at
					 FROM products p
					 WHERE p.id = :product_id`

		query, args, err := sqlx.Named(query, map[string]interface{}{
			"product_id": 4,
		})
		if err != nil {
			log.Fatal(err)
		}

		query = sqlx.Rebind(sqlx.DOLLAR, query)

		p := product{}
		err = db.GetContext(ctx, &p, query, args...)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d — %s (%v)\n", p.ID, p.Name, p.CreatedAt.Year())
	}

	fmt.Println()
	fmt.Println("GetContext + sqlx.In")

	{
		productIDs := []int64{1, 3, 5}

		query := `SELECT p.id, p.name, p.created_at
					 FROM products p
					 WHERE p.id in (:product_ids)`

		query, args, err := sqlx.Named(query, map[string]interface{}{
			"product_ids": productIDs,
		})
		if err != nil {
			log.Fatal(err)
		}

		query, args, err = sqlx.In(query, args...)
		if err != nil {
			log.Fatal(err)
		}

		query = db.Rebind(query) // Можно так

		p := product{}
		err = db.GetContext(ctx, &p, query, args...)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d — %s (%v)\n", p.ID, p.Name, p.CreatedAt.Year())
	}

	fmt.Println()
	fmt.Println("NamedQueryContext")

	{ // NamedQueryContext
		const query = `SELECT p.id, p.name, p.created_at
					 FROM products p
					 WHERE p.category_id = :category_id`

		rows, err := db.NamedQueryContext(ctx, query, map[string]interface{}{
			"category_id": 1,
		})
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var products []product
		for rows.Next() {
			var p product
			err = rows.StructScan(&p)
			if err != nil {
				log.Fatal(err)
			}
			products = append(products, p)
		}

		for _, p := range products {
			fmt.Printf("%d — %s (%v)\n", p.ID, p.Name, p.CreatedAt.Year())
		}
	}

	fmt.Println()
	fmt.Println("SelectContext")

	{
		const query = `SELECT p.id, p.name
					 FROM products p
					 ORDER BY created_at DESC`

		var products []product
		err := db.SelectContext(ctx, &products, query)
		if err != nil {
			log.Fatal(err)
		}

		for _, p := range products {
			fmt.Printf("%d — %s\n", p.ID, p.Name)
		}
	}
}
