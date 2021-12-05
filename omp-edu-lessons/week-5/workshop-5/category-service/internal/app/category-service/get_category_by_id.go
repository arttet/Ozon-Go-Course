package category_service

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	internal_errors "github.com/ozonmp/week-5-workshop/category-service/internal/pkg/errors"
	"github.com/ozonmp/week-5-workshop/category-service/internal/service/category"
	category_service "github.com/ozonmp/week-5-workshop/category-service/pkg/category-service"
)

func (i *Implementation) GetCategoryById(ctx context.Context, req *category_service.GetCategoryByIdRequest) (*category_service.GetCategoryByIdResponse, error) {
	err := validateGetCategoryByIdRequest(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	cat, err := i.categoryService.GetCategoryByID(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, internal_errors.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, errors.Wrap(err, "categoryService.GetCategoryByID")
	}

	return makeGetCategoryByIdResponse(cat), nil
}

func validateGetCategoryByIdRequest(req *category_service.GetCategoryByIdRequest) error {
	return req.Validate()
}

func makeGetCategoryByIdResponse(cat *category.Category) *category_service.GetCategoryByIdResponse {
	pbCat := &category_service.Category{
		Id:   cat.ID,
		Name: cat.Name,
	}

	return &category_service.GetCategoryByIdResponse{
		Category: pbCat,
	}
}
