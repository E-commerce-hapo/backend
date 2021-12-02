package category

import "context"

type ICategoryAggr interface {
	CreateCategory(context.Context, *CreateCategoryArgs) error
}

type ICategoryQuery interface {
	ListCategories(context.Context, *ListCategoriesArgs) ([]*Category, error)
}
