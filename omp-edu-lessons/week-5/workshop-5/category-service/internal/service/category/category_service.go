package category

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"

	internal_errors "github.com/ozonmp/week-5-workshop/category-service/internal/pkg/errors"
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

	// Пример с деградацией

	eg, egCtx := errgroup.WithContext(ctx)

	var categoryIcon string
	eg.Go(func() (egErr error) {
		categoryIcon, egErr = GetCategoryIcon(egCtx, cat.ID)
		return errors.Wrap(egErr, "GetCategoryIcon()")
	})

	var categoryDescription string
	eg.Go(func() (egErr error) {
		categoryDescription, egErr = GetCategoryDescription(egCtx, cat.ID)
		return errors.Wrap(egErr, "GetCategoryIcon()")
	})

	var categoryRating uint8
	eg.Go(func() (egErr error) {
		categoryRating, egErr = GetCategoryRating(egCtx, cat.ID)
		if egErr != nil {
			log.Error().Msg("GetCategoryRating()")
		}
		return nil
	})

	err = eg.Wait()
	if err != nil {
		return nil, err
	}

	// Каким-то образом используем значения
	_, _, _ = categoryIcon, categoryDescription, categoryRating

	return cat, nil
}

// External call stubs:

func GetCategoryIcon(ctx context.Context, catID uint64) (string, error) {
	return "", nil
}

func GetCategoryDescription(ctx context.Context, catID uint64) (string, error) {
	return "", nil
}

func GetCategoryRating(ctx context.Context, catID uint64) (uint8, error) {
	return 0, nil
}
