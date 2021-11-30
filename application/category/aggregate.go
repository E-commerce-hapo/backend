package category

import (
	"context"

	"github.com/kiem-toan/pkg/integration/email"

	"github.com/kiem-toan/pkg/event/dispatcher"

	"github.com/kiem-toan/pkg/idx"

	"github.com/kiem-toan/pkg/database"

	"github.com/kiem-toan/application/category/sqlstore"
	service_category "github.com/kiem-toan/domain/service/category"
)

type CategoryAggregate struct {
	categoryStore sqlstore.CategoryStoreFactory
	dispatcher    *dispatcher.Dispatcher
	emailClient   *email.Client
}

var _ service_category.CategoryAggrService = &CategoryAggregate{}

func NewCategoryAggregate(db *database.Database, dispatcher *dispatcher.Dispatcher, emailClient *email.Client) *CategoryAggregate {
	//productPM *pm.ProductManager
	//defer func() {
	//	err := dispatcher.Register(service_category.CreatedCategoryEventName, []listener.Listener{productPM})
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}()
	return &CategoryAggregate{
		categoryStore: sqlstore.NewCategoryStore(db),
		dispatcher:    dispatcher,
		emailClient:   emailClient,
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
