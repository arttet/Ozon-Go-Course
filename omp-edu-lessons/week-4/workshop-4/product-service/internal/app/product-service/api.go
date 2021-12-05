package product_service

import (
	product_service "github.com/ozonmp/week-4-workshop/product-service/internal/service/product"
	desc "github.com/ozonmp/week-4-workshop/product-service/pkg/product-service"
)

type Implementation struct {
	desc.UnimplementedProductServiceServer
	productService *product_service.Service
}

func NewProductService(productService *product_service.Service) desc.ProductServiceServer {
	return &Implementation{
		productService: productService,
	}
}
