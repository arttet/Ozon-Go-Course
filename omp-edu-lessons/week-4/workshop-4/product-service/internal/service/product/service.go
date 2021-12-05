package product_service

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
	category_service "github.com/ozonmp/week-4-workshop/category-service/pkg/category-service"

)

var ErrWrongCategory = errors.New("category does not exist")

//go:generate mockgen -package=product_service -destination=service_mocks_test.go -self_package=github.com/ozonmp/week-4-workshop/product-service/internal/service/product . IRepository,ICategoryClient

type IRepository interface {
	SaveProduct(ctx context.Context, product *Product) (int64, error)
	DeleteProduct(ctx context.Context, IDs []int64) error
	GetProduct(ctx context.Context, IDs []int64) ([]Product, error)
}

type ICategoryClient interface {
	IsCategoryExists(ctx context.Context, categoryID int64) (ok bool, err error)
}

type Service struct {
	repo   IRepository
	client ICategoryClient
}

func NewService(grpcClient category_service.CategoryServiceClient, db *sqlx.DB) *Service {
	return &Service{
		repo:   newRepo(db),
		client: newClient(grpcClient),
	}
}

func (s *Service) CreateProduct(
	ctx context.Context,
	name string,
	categoryID int64,
	attributes []ProductAttribute,
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
		Attributes: attributes,
	}

	if id, err := s.repo.SaveProduct(ctx, product); err != nil {
		return nil, err
	} else {
		product.ID = id
		return product, nil
	}
}

func (s *Service) DeleteProduct(ctx context.Context, IDs []int64) error {
	return s.repo.DeleteProduct(ctx, IDs)
}
func (s *Service)  GetProduct(ctx context.Context, IDs []int64) ([]Product, error) {
	return s.repo.GetProduct(ctx, IDs)
}
