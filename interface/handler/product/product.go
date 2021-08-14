package product

import (
	"github.com/kiem-toan/interface/controller/product"
)

type ProductHandler struct {
	ProductService *product.ProductService
}

func New(productSvc *product.ProductService) *ProductHandler {
	return &ProductHandler{
		ProductService: productSvc,
	}
}
