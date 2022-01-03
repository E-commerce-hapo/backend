package category

import (
	"time"

	"github.com/E-commerce-hapo/backend/pkg/idx"
	"github.com/E-commerce-hapo/backend/pkg/paging"
)

type Category struct {
	ID          idx.ID
	Name        string
	Description string
	ShopID      idx.ID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type CreateCategoryArgs struct {
	Name        string
	Description string
	ShopID      idx.ID
}

type ListCategoriesArgs struct {
	Paging *paging.Paging
}
