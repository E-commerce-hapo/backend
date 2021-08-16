package category

import (
	"context"
	"time"

	"github.com/k0kubun/pp"

	"github.com/kiem-toan/infrastructure/integration/email"

	"github.com/kiem-toan/infrastructure/event/dispatcher"

	"github.com/kiem-toan/infrastructure/idx"

	"github.com/kiem-toan/infrastructure/database"

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
	event := service_category.CreatedCategoryEvent{
		Time: time.Now().UTC(),
		ID:   "111",
	}
	err := c.dispatcher.Dispatch(event)
	c.emailClient.SendMail(ctx, &email.SendEmailCommand{
		FromName:    "shinichi24567@gmail.com",
		ToAddresses: []string{"1751012015hai@ou.edu.vn"},
		Subject:     "subject",
		Content:     "abcc121",
	})
	if err != nil {
		pp.Println(err)
		return err
	}
	return nil
}
