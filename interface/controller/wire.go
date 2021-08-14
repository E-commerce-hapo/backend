package controller

import (
	"github.com/google/wire"
	"github.com/kiem-toan/interface/controller/category"
	"github.com/kiem-toan/interface/controller/product"
)

var WireSet = wire.NewSet(
	category.New,
	product.New,
)
