package task_repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/ozonmp/week-4-workshop/category-service/internal/service/database"
	taskpkg "github.com/ozonmp/week-4-workshop/category-service/internal/service/task"
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return Repository{db: db}
}

func (r Repository) FindNonStartedTask(ctx context.Context, tx *sqlx.Tx) (*taskpkg.Task, error) {
	sb := database.StatementBuilder.
		Select("id", "started_at").
		From("task").
		Where(sq.Eq{"started_at": nil}).
		Limit(1)

	query, args, err := sb.ToSql()
	if err != nil {
		return nil, err
	}

	var queryer sqlx.QueryerContext
	if tx == nil {
		queryer = r.db
	} else {
		queryer = tx
	}

	task := new(taskpkg.Task)
	err = queryer.QueryRowxContext(ctx, query, args...).StructScan(task)
	if err != nil {
		return nil, errors.Wrap(err, "db.QueryRowxContext()")
	}

	return task, nil
}

func (r Repository) SaveTask(ctx context.Context, task *taskpkg.Task, tx *sqlx.Tx) error {
	sb := database.StatementBuilder.
		Update("task").
		Set("started_at", task.StartedAt).
		Where(sq.And{
			sq.Eq{"id": task.ID},
			sq.Eq{"started_at": nil}, // Лишняя проверка
		})

	query, args, err := sb.ToSql()
	if err != nil {
		return err
	}

	var execer sqlx.ExecerContext
	if tx == nil {
		execer = r.db
	} else {
		execer = tx
	}

	_, err = execer.ExecContext(ctx, query, args...)
	return errors.Wrap(err, "db.ExecContext()")
}
