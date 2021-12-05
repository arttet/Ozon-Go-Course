package database

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

var StatementBuilder = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

func New(ctx context.Context, dsn string) (db *sqlx.DB, err error) {
	db, err = sqlx.Open("pgx", dsn)
	if err != nil {
		err = errors.Wrap(err, "sql.Open()")
		return
	}

	err = db.PingContext(ctx)
	if err != nil {
		err = errors.Wrap(err, "db.PingContext()")
		return
	}
	return
}
