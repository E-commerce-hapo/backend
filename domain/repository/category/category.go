package category

import (
	"context"

	service_category "github.com/kiem-toan/domain/service/category"
)

type CategoryRepositoryService interface {
	CreateCategory(context.Context, *service_category.Category) error
}
