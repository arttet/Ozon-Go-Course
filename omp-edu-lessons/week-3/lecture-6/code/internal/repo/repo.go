package repo

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/ozoncp/ocp-template-api/internal/models"
)

type Repo interface {
	DescribeTemplate(ctx context.Context, templateID uint64) (*models.Template, error)
}

type repo struct {
	db        *sqlx.DB
	batchSize uint
}

func NewRepo(db *sqlx.DB, batchSize uint) Repo {
	return &repo{db: db, batchSize: batchSize}
}

func (r *repo) DescribeTemplate(ctx context.Context, templateID uint64) (*models.Template, error) {
	return nil, nil
}
