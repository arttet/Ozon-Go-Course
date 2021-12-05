package database

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type WithTxFunc func(ctx context.Context, tx *sqlx.Tx) error

func WithTx(ctx context.Context, db *sqlx.DB, fn WithTxFunc) error {
	t, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "db.BeginTxx()")
	}

	err = fn(ctx, t)
	if err != nil {
		if errRollback := t.Rollback(); errRollback != nil {
			return errors.Wrap(err, "Tx.Rollback")
		}
		return errors.Wrap(err, "Tx.WithTxFunc")
	}

	err = t.Commit()
	if err != nil {
		return errors.Wrap(err, "Tx.Commit")
	}
	return nil
}
