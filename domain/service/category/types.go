package category

import (
	"time"

	"github.com/kiem-toan/domain/service/common"

	"github.com/kiem-toan/pkg/idx"
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
	Paging *common.Paging
}
