package product_service

import (
	"context"

	category_service "github.com/ozonmp/week-4-workshop/category-service/pkg/category-service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type categoryClient struct {
	grpcClient category_service.CategoryServiceClient
}

func newClient(grpcClient category_service.CategoryServiceClient) *categoryClient {
	return &categoryClient{
		grpcClient: grpcClient,
	}
}

func (c *categoryClient) IsCategoryExists(ctx context.Context, categoryID int64) (ok bool, err error) {
	_, err = c.grpcClient.GetCategoryById(ctx, &category_service.GetCategoryByIdRequest{
		Id: uint64(categoryID),
	})
	if err == nil {
		return true, nil
	}

	if status.Code(err) == codes.NotFound {
		return false, nil
	}

	return false, err
}
