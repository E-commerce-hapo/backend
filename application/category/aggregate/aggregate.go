package aggregate

import (
	"context"

	"github.com/E-commerce-hapo/backend/application/category"
	"github.com/E-commerce-hapo/backend/application/category/sqlstore"
	"github.com/E-commerce-hapo/backend/pkg/database"
	"github.com/E-commerce-hapo/backend/pkg/event/dispatcher"
	"github.com/E-commerce-hapo/backend/pkg/idx"
	"github.com/E-commerce-hapo/backend/thirdparty/email"
)

type ICategoryAggr interface {
	CreateCategory(context.Context, *category.CreateCategoryArgs) error
}

type categoryAggregate struct {
	categoryStore sqlstore.CategoryStoreFactory
	dispatcher    *dispatcher.Dispatcher
	emailClient   *email.Client
}

var _ ICategoryAggr = &categoryAggregate{}

func NewCategoryAggregate(db *database.Database, dispatcher *dispatcher.Dispatcher, emailClient *email.Client) ICategoryAggr {
	//productPM *pm.ProductManager
	//defer func() {
	//	err := dispatcher.Register(service_category.CreatedCategoryEventName, []listener.Listener{productPM})
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}()
	return &categoryAggregate{
		categoryStore: sqlstore.NewCategoryStore(db),
		dispatcher:    dispatcher,
		emailClient:   emailClient,
	}
}
func (c *categoryAggregate) CreateCategory(ctx context.Context, args *category.CreateCategoryArgs) error {
	category := &category.Category{
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
