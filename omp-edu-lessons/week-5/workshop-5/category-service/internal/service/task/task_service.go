package task

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Service struct {
	repository RepositoryInterface
	db         *sqlx.DB
}

type RepositoryInterface interface {
	FindNonStartedTask(context.Context, *sqlx.Tx) (*Task, error)
	InsertTask(context.Context, *Task, *sqlx.Tx) error
	UpdateTask(context.Context, *Task, *sqlx.Tx) error
	DeleteTask(context.Context, *Task, *sqlx.Tx) error
	GetMetrics(context.Context) (*Metrics, error)
}

func New(repository RepositoryInterface, db *sqlx.DB) Service {
	s := Service{
		repository: repository,
		db:         db,
	}

	s.runTaskProducers()
	s.runTaskConsumers()
	s.runTaskMetrics()

	return s
}
