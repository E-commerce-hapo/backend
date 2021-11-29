package category

import (
	"context"

	"github.com/kiem-toan/application/category/sqlstore"
	service_category "github.com/kiem-toan/domain/service/category"
	"github.com/kiem-toan/pkg/database"
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

func (c CategoryQuery) ListCategories(ctx context.Context, args *service_category.ListCategoriesArgs) ([]*service_category.Category, error) {
	query, err := c.categoryStore(ctx).WithPaging(ctx, args.Paging)
	if err != nil {
		return nil, err
	}
	categories, err := query.ListCategories(ctx)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
