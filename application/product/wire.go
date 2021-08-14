package product

import (
	"github.com/google/wire"
	"github.com/kiem-toan/application/product/pm"
)

var WireSet = wire.NewSet(
	pm.New, NewProductAggregate,
)
