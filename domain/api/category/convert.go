package category

import (
	service_category "github.com/kiem-toan/domain/service/category"
)

func Convert_service_Category_to_api_Category(in *service_category.Category) *Category {
	if in == nil {
		return nil
	}
	out := &Category{
		ID:          in.ID,
		Name:        in.Name,
		Description: in.Description,
		UpdatedAt:   in.UpdatedAt,
		CreatedAt:   in.CreatedAt,
	}
	return out
}

func Convert_service_Categories_to_api_Categories(in []*service_category.Category) []*Category {
	var out []*Category
	for _, cateModel := range in {
		cate := Convert_service_Category_to_api_Category(cateModel)
		out = append(out, cate)
	}
	return out
}
