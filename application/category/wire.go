package category

import "github.com/google/wire"

var WireSet = wire.NewSet(
	NewCategoryAggregate, NewCategoryQuery,
)
