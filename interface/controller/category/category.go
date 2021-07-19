package category

import (
	"github.com/kiem-toan/domain/api/category"
	"github.com/kiem-toan/infrastructure/database"
	"github.com/kiem-toan/infrastructure/errorx"
	"gorm.io/gorm"
	"net/http"
)

type CategoryService struct {
	db *database.Database
}

var _ category.ICategoryService = &CategoryService{}

func New(db *database.Database) *CategoryService {
	return &CategoryService{
		db: db,
	}
}

type ProductVariant struct {
	gorm.Model
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (*ProductVariant) TableName() string {
	return "product_variant"
}


func (t *CategoryService) CreateCategory(i *category.CreateCategoryRequest) (*category.CreateCategoryResponse, error){
	if i.ID == "1" {
		return &category.CreateCategoryResponse{
			ID:   "11111",
			Name: "abccc",
		}, nil
	}
	return nil, errorx.New(http.StatusBadRequest, nil, "bad request nè")
}

