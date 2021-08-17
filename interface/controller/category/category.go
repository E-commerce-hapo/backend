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

func (t *CategoryService) ListCategories(ctx context.Context, r *category.ListCategoriesRequest) (*category.ListCategoriesResponse, error) {
	svcPaging, err := r.Paging.Convert_api_Paging_to_service_Paging()
	if err != nil {
		return nil, err
	}
	args := &service_category.ListCategoriesArgs{
		Paging: svcPaging,
	}
	categories, err := t.CategoryQuery.ListCategories(ctx, args)
	if err != nil {
		return nil, err
	}

	return &category.ListCategoriesResponse{
		Categories: category.Convert_service_Categories_to_api_Categories(categories),
		Paging:     r.Paging.Convert_api_Paging_to_api_PagingInfo(len(categories)),
	}, nil
}
