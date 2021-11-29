package category

import (
	"github.com/kiem-toan/pkg/idx"
	"gorm.io/gorm"
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
