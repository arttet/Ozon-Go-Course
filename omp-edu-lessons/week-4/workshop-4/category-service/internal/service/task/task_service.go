package task

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/ozonmp/week-4-workshop/category-service/internal/service/database"
)

type Service struct {
	repository RepositoryInterface
	db         *sqlx.DB
}

type RepositoryInterface interface {
	FindNonStartedTask(context.Context, *sqlx.Tx) (*Task, error)
	SaveTask(context.Context, *Task, *sqlx.Tx) error
}

func New(repository RepositoryInterface, db *sqlx.DB) Service {
	return Service{
		repository: repository,
		db:         db,
	}
}

func (s Service) ExecTask(ctx context.Context) error {
	txErr := database.WithTx(ctx, s.db, func(ctx context.Context, tx *sqlx.Tx) error {
		task, err := s.repository.FindNonStartedTask(ctx, tx)
		if err != nil {
			return errors.Wrap(err, "repository.FindNonStartedTask()")
		}

		now := time.Now()
		task.StartedAt = &now

		err = s.repository.SaveTask(ctx, task, tx)
		if err != nil {
			return errors.Wrap(err, "repository.SaveTask()")
		}
		return nil
	})
	if txErr != nil {
		return txErr
	}

	time.Sleep(time.Second) // Имитация работы

	return nil
}
