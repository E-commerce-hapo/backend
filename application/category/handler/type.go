package handler

import (
	"time"

	"github.com/E-commerce-hapo/backend/pkg/idx"
	"github.com/E-commerce-hapo/backend/pkg/paging"
)

type Category struct {
	ID          idx.ID    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateCategoryRequest struct {
	// ID tự gen ra ở server
	ID string `json:"id"`
	// Tên của danh mục, bắt buộc phải có
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

// swagger:response CreateCategoryResponse
type CreateCategoryResponse struct {
	ID   string `json:"id"`
	Name string `json:"name "`
}

type ListCategoriesRequest struct {
	Paging *paging.Paging `json:"paging"`
}

type ListCategoriesResponse struct {
	Categories []*Category    `json:"categories"`
	Paging     *paging.Paging `json:"paging"`
}
