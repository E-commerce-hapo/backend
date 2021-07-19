package category

import (
	"context"

	"github.com/kiem-toan/infrastructure/idx"

	"github.com/kiem-toan/infrastructure/database"

	"github.com/kiem-toan/application/category/sqlstore"
	service_category "github.com/kiem-toan/domain/service/category"
)

type CategoryAggregate struct {
	categoryStore sqlstore.CategoryStoreFactory
}

var _ service_category.CategoryService = &CategoryAggregate{}

func NewCategoryAggregate(db *database.Database) *CategoryAggregate {
	return &CategoryAggregate{
		categoryStore: sqlstore.NewCategoryStore(db),
	}
}
func (c *CategoryAggregate) CreateCategory(ctx context.Context, args *service_category.CreateCategoryArgs) error {
	category := &service_category.Category{
		ID:          idx.NewID(),
		Name:        args.Name,
		Description: args.Description,
		ShopID:      args.ShopID,
	}
	if err := c.categoryStore(ctx).CreateCategory(ctx, category); err != nil {
		return err
	}
	return nil
}
