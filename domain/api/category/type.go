package category

import (
	"time"

	"github.com/kiem-toan/infrastructure/idx"
)

type Category struct {
	ID          idx.ID    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
	DeletedAt   time.Time `json:"deleted_at"`
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