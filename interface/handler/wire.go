package handler

import (
	"github.com/google/wire"
	"github.com/kiem-toan/interface/handler/category"
	"github.com/kiem-toan/interface/handler/product"
)

var WireSet = wire.NewSet(
	category.New, product.New,
)
