package category

import "context"

type CategoryAggrService interface {
	CreateCategory(context.Context, *CreateCategoryArgs) error
}

type CategoryQueryService interface {
	ListCategories(context.Context, *CreateCategoryArgs) ([]*Category, error)
}
