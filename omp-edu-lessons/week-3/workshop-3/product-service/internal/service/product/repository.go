package product_service

import "context"

var nextID int64 = 1

type repository struct{}

func newRepo() IRepository {
	return repository{}
}

func (r repository) SaveProduct(ctx context.Context, product *Product) error {
	product.ID = nextID

	nextID++

	return nil
}
