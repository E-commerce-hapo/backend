package query

import (
	"context"

	"github.com/E-commerce-hapo/backend/application/category"

	"github.com/E-commerce-hapo/backend/application/category/sqlstore"
	"github.com/E-commerce-hapo/backend/pkg/database"
)

type ICategoryQuery interface {
	ListCategories(context.Context, *category.ListCategoriesArgs) ([]*category.Category, error)
}

type categoryQuery struct {
	categoryStore sqlstore.CategoryStoreFactory
}

var _ ICategoryQuery = &categoryQuery{}

func NewCategoryQuery(db *database.Database) *categoryQuery {
	return &categoryQuery{
		categoryStore: sqlstore.NewCategoryStore(db),
	}
}

func (c categoryQuery) ListCategories(ctx context.Context, args *category.ListCategoriesArgs) ([]*category.Category, error) {
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
