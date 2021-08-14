package product

import (
	"context"

	"github.com/kiem-toan/domain/api/product"
)

type ProductService struct {
}

var _ product.IProductService = &ProductService{}

func New() *ProductService {
	return &ProductService{}
}
func (p ProductService) CreateProduct(ctx context.Context) error {
	panic("implement me")
}
