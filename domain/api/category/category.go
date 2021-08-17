package category

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/kiem-toan/infrastructure/httpx"
)

type ICategoryService interface {
	CreateCategory(context.Context, *CreateCategoryRequest) (*httpx.CreatedResponse, error)
	ListCategories(*gin.Context, *ListCategoriesRequest) (*ListCategoriesResponse, error)
}
