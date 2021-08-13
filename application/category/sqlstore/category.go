package sqlstore

import (
	"context"
	"fmt"

	"github.com/k0kubun/pp"

	"github.com/kiem-toan/infrastructure/idx"

	service_category "github.com/kiem-toan/domain/service/category"

	repo_cateogry "github.com/kiem-toan/domain/repository/category"

	"gorm.io/gorm"

	"github.com/kiem-toan/infrastructure/database"
)

type CategoryStore struct {
	gormDB *gorm.DB
	preds  []string
}

var _ repo_cateogry.CategoryRepositoryService = &CategoryStore{}

type CategoryStoreFactory func(ctx context.Context) *CategoryStore

func NewCategoryStore(db *database.Database) CategoryStoreFactory {
	return func(ctx context.Context) *CategoryStore {
		return &CategoryStore{
			gormDB: db.Db,
		}
	}
}

func (s *CategoryStore) CreateCategory(ctx context.Context, category *service_category.Category) error {
	categoryDB := &repo_cateogry.Category{
		ID:          idx.NewID(),
		Name:        category.Name,
		Description: category.Description,
		ShopID:      category.ShopID,
	}
	return s.createCategoryDB(ctx, categoryDB)
}

func (s *CategoryStore) createCategoryDB(ctx context.Context, categoryDB *repo_cateogry.Category) error {
	tx := s.gormDB.Create(categoryDB)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (s *CategoryStore) IsDeleted() *CategoryStore {
	s.preds = append(s.preds, "deleted_at IS NULL ")
	return s
}

func (s *CategoryStore) ByName(name string) *CategoryStore {
	s.preds = append(s.preds, fmt.Sprintf("name = '%v'", name))
	return s
}

func (s *CategoryStore) ListCategoriesDB(ctx context.Context) (categoriesDB []*repo_cateogry.Category, err error) {
	pp.Println(s.preds)
	//tx := s.gormDB.Where(strings.Join(s.preds, "AND")).Find(&categoriesDB)
	tx := s.gormDB.Where("name = 'abcc'").Find(&categoriesDB)
	pp.Println(categoriesDB)
	pp.Println(len(categoriesDB))
	if tx.Error != nil {
		return nil, tx.Error
	}
	return categoriesDB, nil
}

func (s *CategoryStore) ListCategories(ctx context.Context) ([]*service_category.Category, error) {
	categories, err := s.ListCategoriesDB(ctx)
	if err != nil {
		return nil, err
	}
	return repo_cateogry.Convert_model_Categories_to_service_Categories(categories), nil
}
