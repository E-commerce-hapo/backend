package product

import "context"

type ProductAggrService interface {
	CreateProduct(context.Context) error
}
