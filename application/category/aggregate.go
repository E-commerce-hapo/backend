package category

import (
	"context"
	"time"

	"github.com/kiem-toan/infrastructure/event/dispatcher"

	"github.com/kiem-toan/infrastructure/idx"

	"github.com/kiem-toan/infrastructure/database"

	"github.com/kiem-toan/application/category/sqlstore"
	service_category "github.com/kiem-toan/domain/service/category"
)

type CategoryAggregate struct {
	categoryStore sqlstore.CategoryStoreFactory
	dispatcher    *dispatcher.Dispatcher
}

var _ service_category.CategoryAggrService = &CategoryAggregate{}

func NewCategoryAggregate(db *database.Database, dispatcher *dispatcher.Dispatcher) *CategoryAggregate {
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
	event := service_category.CreatedCategoryEvent{
		Time: time.Now().UTC(),
		ID:   "111",
	}
	err := c.dispatcher.Dispatch(event)
	if err != nil {
		return err
	}
	return nil
}
