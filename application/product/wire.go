package product

import (
	"github.com/E-commerce-hapo/backend/application/product/pm"
	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	pm.New, NewProductAggregate,
)
