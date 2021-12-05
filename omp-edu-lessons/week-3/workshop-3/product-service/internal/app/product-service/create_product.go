package product_service

import (
	"context"

	product_service "github.com/ozonmp/week-3-workshop/product-service/internal/service/product"
	desc "github.com/ozonmp/week-3-workshop/product-service/pkg/product-service"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) CreateProduct(ctx context.Context, req *desc.CreateProductRequest) (*desc.CreateProductResponse, error) {
	res, err := i.productService.CreateProduct(ctx, req.GetName(), req.GetCategoryId())
	if err != nil {
		if err == product_service.ErrWrongCategory {
			details := &errdetails.BadRequest{
				FieldViolations: []*errdetails.BadRequest_FieldViolation{
					{
						Field:       "categoryId",
						Description: "wrong category",
					},
				},
			}

			st := status.New(codes.InvalidArgument, "wrong category")

			withDetails, stErr := st.WithDetails(details)
			if stErr != nil {
				return nil, st.Err()
			}

			return nil, withDetails.Err()
		}

		return nil, err
	}

	return &desc.CreateProductResponse{
		Result: convertProductToPb(res),
	}, nil
}

func convertProductToPb(res *product_service.Product) *desc.Product {
	return &desc.Product{
		Id:         res.ID,
		Name:       res.Name,
		CategoryId: res.CategoryID,
	}
}
