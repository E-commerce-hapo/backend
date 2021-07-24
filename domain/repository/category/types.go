package category

import (
	"time"

	"github.com/kiem-toan/infrastructure/idx"
	"gorm.io/gorm"
)

// gorm.Model
type Category struct {
	gorm.Model
	ID          idx.ID `gorm:"primaryKey"`
	Name        string `gorm:"name"`
	Description string
	ShopID      idx.ID
	CreatedAt   time.Time
	DeletedAt   time.Time
	UpdatedAt   time.Time
}

func (*Category) TableName() string {
	return "category"
}
