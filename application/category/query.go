package category

import (
	"context"

	"github.com/kiem-toan/application/category/sqlstore"
	service_category "github.com/kiem-toan/domain/service/category"
	"github.com/kiem-toan/infrastructure/database"
)

type CategoryQuery struct {
	categoryStore sqlstore.CategoryStoreFactory
}

var _ service_category.CategoryQueryService = &CategoryQuery{}

func NewCategoryQuery(db *database.Database) *CategoryQuery {
	return &CategoryQuery{
		categoryStore: sqlstore.NewCategoryStore(db),
	}
}

func (c CategoryQuery) ListCategories(ctx context.Context, args *service_category.CreateCategoryArgs) ([]*service_category.Category, error) {
	categories, err := c.categoryStore(ctx).ListCategories(ctx)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
