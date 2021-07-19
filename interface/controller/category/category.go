package category

import (
	"context"

	"github.com/kiem-toan/infrastructure/httpx"

	"github.com/kiem-toan/infrastructure/idx"

	service_category "github.com/kiem-toan/domain/service/category"

	categorying "github.com/kiem-toan/application/category"

	"github.com/kiem-toan/domain/api/category"
)

type CategoryService struct {
	CategoryAgg *categorying.CategoryAggregate
}

var _ category.ICategoryService = &CategoryService{}

func New(cateAgg *categorying.CategoryAggregate) *CategoryService {
	return &CategoryService{
		CategoryAgg: cateAgg,
	}
}

func (t *CategoryService) CreateCategory(ctx context.Context, r *category.CreateCategoryRequest) (*httpx.CreatedResponse, error) {
	category := &service_category.CreateCategoryArgs{
		Name:        r.Name,
		Description: r.Description,
		ShopID:      idx.NewID(),
	}
	if err := t.CategoryAgg.CreateCategory(ctx, category); err != nil {
		return nil, err
	}
	return &httpx.CreatedResponse{
		Created: 1,
	}, nil
}
