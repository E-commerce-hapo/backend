package product

import (
	"context"

	"github.com/E-commerce-hapo/backend/pkg/event/dispatcher"
)

type ProductAggregate struct {
	dispatcher *dispatcher.Dispatcher
}

//var _ service_category.ProductAggrService = &ProductAggregate{}

func NewProductAggregate(dispatcher *dispatcher.Dispatcher) *ProductAggregate {
	return &ProductAggregate{
		dispatcher: dispatcher,
	}
}
func (p ProductAggregate) CreateProduct(ctx context.Context) error {
	panic("implement me")
}
