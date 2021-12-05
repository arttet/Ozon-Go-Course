package product_service

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
)

type repository struct {
	DB *sqlx.DB
}

func (r repository) DeleteProduct(ctx context.Context, IDs []int64) error {
	query := sq.Delete("products").PlaceholderFormat(sq.Dollar).Where(sq.Eq{"id": IDs})
	s, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = r.DB.ExecContext(ctx, s, args...)
	return err
}

func (r repository) GetProduct(ctx context.Context, IDs []int64) ([]Product, error) {
	query := sq.Select("*").PlaceholderFormat(sq.Dollar).From("products").Where(sq.Eq{"id": IDs})

	s, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	var res []Product
	err = r.DB.SelectContext(ctx, &res, s, args...)

	return res, err
}

func newRepo(db *sqlx.DB) IRepository {
	return repository{
		DB: db,
	}
}

func (r repository) SaveProduct(ctx context.Context, product *Product) (int64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "repo.SaveProduct")
	defer span.Finish()

	query := sq.Insert("products").PlaceholderFormat(sq.Dollar).Columns(
		"name", "category_id", "info").Values(product.Name, product.CategoryID, product.Attributes).Suffix("RETURNING id").RunWith(r.DB)

	rows, err := query.QueryContext(ctx)
	if err != nil {
		return 0, err
	}

	var id int64
	if rows.Next() {
		err = rows.Scan(&id)

		if err != nil {
			return 0, err
		}

		return id, nil
	} else {
		return 0, sql.ErrNoRows
	}
}
