package registry

import (
	"github.com/E-commerce-hapo/backend/application/category/aggregate"
	"github.com/E-commerce-hapo/backend/application/category/query"
)

// Category Aggr ...
func (r *Registry) RegisterCategoryAggr() aggregate.ICategoryAggr {
	return aggregate.NewCategoryAggregate(r.DB, r.Dispatcher, r.EmailClient)
}

// Category Query ...
func (r *Registry) RegisterCategoryQuery() query.ICategoryQuery {
	return query.NewCategoryQuery(r.DB)
}
