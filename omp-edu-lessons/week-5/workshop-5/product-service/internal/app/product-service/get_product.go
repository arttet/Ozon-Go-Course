package product_service

import (
	"context"

	desc "github.com/ozonmp/week-5-workshop/product-service/pkg/product-service"
)

func (i *Implementation) GetProduct(ctx context.Context, req *desc.GetProductRequest) (*desc.GetProductResponse, error) {
	products, err := i.productService.GetProduct(ctx, req.GetProductIds())
	if err != nil {
		return nil, err
	}

	pr := make([]*desc.Product, len(products))
	for idx := range products {
		pr[idx] = convertProductToPb(&products[idx])
	}

	return &desc.GetProductResponse{
		Products: pr,
	}, nil
}
