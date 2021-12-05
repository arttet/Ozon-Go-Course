package product_service

import (
	"context"
	"errors"

	category_service "github.com/ozonmp/week-3-workshop/category-service/pkg/category-service"
)

var ErrWrongCategory = errors.New("category does not exist")

//go:generate mockgen -package=product_service -destination=service_mocks_test.go -self_package=github.com/ozonmp/week-3-workshop/product-service/internal/service/product . IRepository,ICategoryClient

type IRepository interface {
	SaveProduct(ctx context.Context, product *Product) error
}

type ICategoryClient interface {
	IsCategoryExists(ctx context.Context, categoryID int64) (ok bool, err error)
}

type Service struct {
	repo   IRepository
	client ICategoryClient
}

func NewService(grpcClient category_service.CategoryServiceClient) *Service {
	return &Service{
		repo:   newRepo(),
		client: newClient(grpcClient),
	}
}

func (s *Service) CreateProduct(
	ctx context.Context,
	name string,
	categoryID int64,
) (*Product, error) {
	exists, err := s.client.IsCategoryExists(ctx, categoryID)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, ErrWrongCategory
	}

	product := &Product{
		Name:       name,
		CategoryID: categoryID,
	}

	if err := s.repo.SaveProduct(ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}
