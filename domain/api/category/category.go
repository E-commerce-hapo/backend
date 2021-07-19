package category

import (
	"context"

	"github.com/kiem-toan/infrastructure/httpx"
)

type ICategoryService interface {
	CreateCategory(ctx context.Context, test *CreateCategoryRequest) (*httpx.CreatedResponse, error)
}

// swagger:parameters Category CreateCategoryRequest
type CreateCategoryRequest struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// swagger:response CreateCategoryResponse
type CreateCategoryResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
