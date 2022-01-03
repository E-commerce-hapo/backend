package sqlstore

import (
	"context"
	"strings"

	"github.com/E-commerce-hapo/backend/pkg/errorx"

	"github.com/E-commerce-hapo/backend/application/category"
	"github.com/E-commerce-hapo/backend/entity"

	"github.com/E-commerce-hapo/backend/pkg/paging"

	"github.com/E-commerce-hapo/backend/pkg/idx"

	"gorm.io/gorm"

	"github.com/E-commerce-hapo/backend/pkg/database"
)

var ErrAgencyNotFound = errorx.Errorf(404, nil, "agency_not_found")

type ICategoryRepository interface {
	CreateCategory(context.Context, *category.Category) error
	ListCategories(context.Context) ([]*category.Category, error)
}

type categoryRepository struct {
	gormDB *gorm.DB
}

var _ ICategoryRepository = &categoryRepository{}

type CategoryStoreFactory func(ctx context.Context) *categoryRepository

func NewCategoryStore(db *database.Database) CategoryStoreFactory {
	return func(ctx context.Context) *categoryRepository {
		return &categoryRepository{
			gormDB: db.Db,
		}
	}
}

func (s *categoryRepository) WithPaging(ctx context.Context, p *paging.Paging) (*categoryRepository, error) {
	err := p.Validate(&entity.Category{})
	if err != nil {
		return nil, err
	}

	s.gormDB = s.gormDB.Limit(p.Limit).Offset(p.Offset)
	if len(p.Sorts) > 0 {
		sort := strings.Join(p.Sorts, ",")
		s.gormDB = s.gormDB.Order(sort)
	}
	return s, nil
}

func (s *categoryRepository) CreateCategory(ctx context.Context, category *category.Category) error {
	s.gormDB.AutoMigrate(&entity.Category{})
	categoryDB := &entity.Category{
		ID:          idx.NewID(),
		Name:        category.Name,
		Description: category.Description,
		ShopID:      category.ShopID,
	}
	return s.createCategoryDB(ctx, categoryDB)
}

func (s *categoryRepository) createCategoryDB(ctx context.Context, categoryDB *entity.Category) error {
	tx := s.gormDB.Create(categoryDB)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (s *categoryRepository) listCategoriesDB(ctx context.Context) (categoriesDB []*entity.Category, err error) {
	tx := s.gormDB.Find(&categoriesDB)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return categoriesDB, nil
}

func (s *categoryRepository) ListCategories(ctx context.Context) ([]*category.Category, error) {
	s.gormDB.AutoMigrate(&entity.Category{})
	categories, err := s.listCategoriesDB(ctx)
	if err != nil {
		return nil, err
	}
	return entity.Convert_model_Categories_to_service_Categories(categories), nil
}
