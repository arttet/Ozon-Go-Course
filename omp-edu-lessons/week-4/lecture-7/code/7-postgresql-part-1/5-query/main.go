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

	db, openErr := sql.Open("pgx", "postgres://user:password@localhost:5432/db")
	if openErr != nil {
		log.Fatalf("sql.Open(): %v", openErr)
	}

	ctx := context.Background()

	{ // Выполнение запросов
		const query = `INSERT INTO products(name, category_id, created_at)
					   VALUES($1, $2, now())
					   RETURNING id`

		result, err := db.ExecContext(ctx, query, "Beret", 4)
		if err != nil {
			log.Fatal(err)
		}

		// result.LastInsertId() - не поддерживается

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("rowsAffected", rowsAffected)
	}

	fmt.Println()

	{ // Получение одной строки
		const query = `SELECT p.id, p.name FROM products p WHERE p.id = $1`

		row := db.QueryRowContext(ctx, query, 2)
		var (
			id   int64
			name string
		)
		err := row.Scan(&id, &name)
		if err == sql.ErrNoRows {
			fmt.Println("row is not found")
		} else if err != nil {
			log.Fatal(err)
		}

		fmt.Println("id", id, "name", name)
	}

	fmt.Println()

	{ // Выполнение запроса с возвращаемым значением
		const query = `INSERT INTO products(name, category_id, created_at)
					   VALUES($1, $2, now())
					   RETURNING id`

		row := db.QueryRowContext(ctx, query, "Ascot cap", 4)

		var id int64
		err := row.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("new record id", id)
	}

	fmt.Println()

	{ // Получение нескольких строк
		const query = `SELECT p.id, p.name FROM products p WHERE p.category_id = $1`

		rows, err := db.QueryContext(ctx, query, 1)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		type product struct {
			id   int64
			name string
		}

		var products []product
		for rows.Next() {
			var p product
			err = rows.Scan(&p.id, &p.name)
			if err != nil {
				log.Fatal(err)
			}
			products = append(products, p)
		}

		for _, p := range products {
			fmt.Printf("%d — %s\n", p.id, p.name)
		}
	}
}
