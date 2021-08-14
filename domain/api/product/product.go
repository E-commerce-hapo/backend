package product

import (
	"context"
)

type IProductService interface {
	CreateProduct(context.Context) error
}
