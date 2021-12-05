package cat_repository

import (
	"context"

	"github.com/ozonmp/week-3-workshop/category-service/internal/service/category"
)

var categories = category.Categories{
	{
		ID:   1,
		Name: "Toys",
	},
	{
		ID:   2,
		Name: "Laptops",
	},
	{
		ID:   3,
		Name: "Auto",
	},
}

type Repository struct{}

func New( /* db */) *Repository {
	return &Repository{}
}

func (Repository) FindAllCategories(_ context.Context) (category.Categories, error) {
	return categories, nil
}
