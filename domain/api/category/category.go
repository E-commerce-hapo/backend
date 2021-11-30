package category

import (
	"context"

	"github.com/kiem-toan/pkg/httpx"
)

type ICategoryService interface {
	CreateCategory(context.Context, *CreateCategoryRequest) (*httpx.CreatedResponse, error)
	ListCategories(context.Context, *ListCategoriesRequest) (*ListCategoriesResponse, error)
}
