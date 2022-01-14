package entity

import (
	"gorm.io/gorm"

	"github.com/E-commerce-hapo/backend/application/category"
	"github.com/E-commerce-hapo/backend/pkg/idx"
)

// gorm.Model
type Category struct {
	gorm.Model
	ID          idx.ID `gorm:"primaryKey"`
	Name        string `gorm:"name"`
	Description string
	ShopID      idx.ID
}

func (*Category) TableName() string {
	return "category"
}

func Convert_model_Category_to_service_Category(in *Category) *category.Category {
	if in == nil {
		return nil
	}
	out := &category.Category{
		ID:          in.ID,
		Name:        in.Name,
		Description: in.Description,
		CreatedAt:   in.CreatedAt,
		UpdatedAt:   in.UpdatedAt,
	}
	return out
}

func Convert_model_Categories_to_service_Categories(in []*Category) []*category.Category {
	var out []*category.Category
	for _, cateModel := range in {
		cate := Convert_model_Category_to_service_Category(cateModel)
		out = append(out, cate)
	}
	return out
}
