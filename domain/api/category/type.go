package category

import (
	"time"

	"github.com/kiem-toan/pkg/idx"
)

type Category struct {
	ID          idx.ID    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}
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

type ListCategoriesRequest struct {
	Categories []*Category `json:"categories"`
}
