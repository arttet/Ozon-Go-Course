package category

import (
	"context"

	"github.com/pkg/errors"

	internal_errors "github.com/ozonmp/week-4-workshop/category-service/internal/pkg/errors"
)

type Service struct {
	repository RepositoryInterface
}

type RepositoryInterface interface {
	FindAllCategories(context.Context) (Categories, error)
}

func New(repository RepositoryInterface) Service {
	return Service{
		repository: repository,
	}
}

var ErrNoCategory = errors.Wrap(internal_errors.ErrNotFound, "no category")

func (s Service) GetCategoryByID(ctx context.Context, id uint64) (*Category, error) {
	cats, err := s.repository.FindAllCategories(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "repository.FindAllCategories")
	}

	cat := cats.FilterByID(id)
	if cat == nil {
		return nil, ErrNoCategory
	}

	return cat, nil
}
