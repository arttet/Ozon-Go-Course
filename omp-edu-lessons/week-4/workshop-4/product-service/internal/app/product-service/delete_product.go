package product_service

import (
	"context"


	desc "github.com/ozonmp/week-4-workshop/product-service/pkg/product-service"
)


func (i *Implementation)  DeleteProduct(ctx context.Context, req *desc.DeleteProductRequest) (*desc.DeleteProductResponse, error) {
	err := i.productService.DeleteProduct(ctx, req.GetProductIds())
	if err != nil {
		return nil, err
	}

	return &desc.DeleteProductResponse{}, nil
}
