package category

import (
	"context"

	"github.com/kiem-toan/pkg/httpx"

	"github.com/kiem-toan/pkg/idx"

	service_category "github.com/kiem-toan/domain/service/category"

	categorying "github.com/kiem-toan/application/category"

	"github.com/kiem-toan/domain/api/category"
)

type CategoryService struct {
	CategoryAgg   *categorying.CategoryAggregate
	CategoryQuery *categorying.CategoryQuery
}

var _ category.ICategoryService = &CategoryService{}

func New(cateAgg *categorying.CategoryAggregate, cateQuery *categorying.CategoryQuery) *CategoryService {
	return &CategoryService{
		CategoryAgg:   cateAgg,
		CategoryQuery: cateQuery,
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
