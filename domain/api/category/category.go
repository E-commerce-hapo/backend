package category

import (
	"context"

	"github.com/kiem-toan/infrastructure/httpx"
)

type ICategoryService interface {
	CreateCategory(context.Context, *CreateCategoryRequest) (*httpx.CreatedResponse, error)
	ListCategories(context.Context, *CreateCategoryRequest) (*ListCategoriesRequest, error)
}
