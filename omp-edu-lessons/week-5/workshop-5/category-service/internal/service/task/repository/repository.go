package task_repository

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/ozonmp/week-5-workshop/category-service/internal/service/database"
	taskpkg "github.com/ozonmp/week-5-workshop/category-service/internal/service/task"
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return Repository{db: db}
}

func (r Repository) FindNonStartedTask(ctx context.Context, tx *sqlx.Tx) (*taskpkg.Task, error) {
	startedGraceTime := time.Now().Add(-time.Minute)

	sb := database.StatementBuilder.
		Select("t.id", "t.exec_duration", "t.started_at").
		From("task t").
		Where(sq.Or{
			sq.Eq{"started_at": nil},
			sq.Lt{"started_at": startedGraceTime},
		}).
		Suffix("FOR UPDATE SKIP LOCKED").
		OrderBy("id").
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
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "db.QueryRowxContext()")
	}

	return task, nil
}

func (r Repository) InsertTask(ctx context.Context, task *taskpkg.Task, tx *sqlx.Tx) error {
	sb := database.StatementBuilder.
		Insert("task").
		SetMap(map[string]interface{}{
			"exec_duration": task.ExecDuration,
			"started_at":    task.StartedAt,
		}).
		Suffix("RETURNING id")

	query, args, err := sb.ToSql()
	if err != nil {
		return err
	}

	var queryer sqlx.QueryerContext
	if tx == nil {
		queryer = r.db
	} else {
		queryer = tx
	}

	err = queryer.QueryRowxContext(ctx, query, args...).Scan(&task.ID)
	return errors.Wrap(err, "db.QueryRowxContext()")
}

func (r Repository) UpdateTask(ctx context.Context, task *taskpkg.Task, tx *sqlx.Tx) error {
	sb := database.StatementBuilder.
		Update("task").
		Where(sq.Eq{"id": task.ID}).
		SetMap(map[string]interface{}{
			"exec_duration": task.ExecDuration,
			"started_at":    task.StartedAt,
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

func (r Repository) DeleteTask(ctx context.Context, task *taskpkg.Task, tx *sqlx.Tx) error {
	sb := database.StatementBuilder.
		Delete("task").
		Where(sq.Eq{"id": task.ID})

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
