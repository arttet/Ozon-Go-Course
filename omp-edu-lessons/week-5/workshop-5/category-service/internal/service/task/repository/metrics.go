package task_repository

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"

	"github.com/ozonmp/week-5-workshop/category-service/internal/service/database"
	taskpkg "github.com/ozonmp/week-5-workshop/category-service/internal/service/task"
)

func (r Repository) GetMetrics(ctx context.Context) (*taskpkg.Metrics, error) {
	sb := database.StatementBuilder.
		Select(
			"count(1) as all_count",
			"count(1) filter (where started_at is null) as non_started_count",
			"count(1) filter (where started_at is not null) as started_count",
		).
		From("task")

	query, args, err := sb.ToSql()
	if err != nil {
		return nil, err
	}

	metrics := new(taskpkg.Metrics)
	err = r.db.QueryRowxContext(ctx, query, args...).StructScan(metrics)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "db.QueryRowxContext()")
	}
	return metrics, nil
}
