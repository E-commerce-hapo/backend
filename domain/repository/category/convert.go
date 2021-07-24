package category

import (
	service_category "github.com/kiem-toan/domain/service/category"
)

func Convert_model_Category_to_service_Category(in *Category) *service_category.Category {
	if in == nil {
		return nil
	}
	out := &service_category.Category{
		ID:          in.ID,
		Name:        in.Name,
		Description: in.Description,
		CreatedAt:   in.CreatedAt,
		UpdatedAt:   in.UpdatedAt,
	}
	return out
}

func Convert_model_Categories_to_service_Categories(in []*Category) []*service_category.Category {
	var out []*service_category.Category
	for _, cateModel := range in {
		cate := Convert_model_Category_to_service_Category(cateModel)
		out = append(out, cate)
	}
	return out
}
