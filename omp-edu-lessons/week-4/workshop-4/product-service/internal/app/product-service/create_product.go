package product_service

import (
	"context"

	product_service "github.com/ozonmp/week-4-workshop/product-service/internal/service/product"
	desc "github.com/ozonmp/week-4-workshop/product-service/pkg/product-service"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) CreateProduct(ctx context.Context, req *desc.CreateProductRequest) (*desc.CreateProductResponse, error) {
	attributes := make([]product_service.ProductAttribute, len(req.GetAttributes()))
	for idx, val := range req.GetAttributes() {
		attributes[idx] = convertPbToProductAttributes(val)
	}

	res, err := i.productService.CreateProduct(ctx, req.GetName(), req.GetCategoryId(), attributes)
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
	resAttr := make([]*desc.ProductAttribute, len(res.Attributes))
	for idx, val := range res.Attributes {
		resAttr[idx] = convertProductAttributesToPb(val)
	}
	return &desc.Product{
		Id:         res.ID,
		Name:       res.Name,
		CategoryId: res.CategoryID,
		Attributes: resAttr,
	}
}

func convertPbToProductAttributes(pa *desc.ProductAttribute) product_service.ProductAttribute {
	return product_service.ProductAttribute{
		ID:    pa.GetId(),
		Value: pa.GetValue(),
	}
}

func convertProductAttributesToPb(pa product_service.ProductAttribute ) *desc.ProductAttribute {
	return &desc.ProductAttribute{
		Id:    pa.ID,
		Value: pa.Value,
	}
}
