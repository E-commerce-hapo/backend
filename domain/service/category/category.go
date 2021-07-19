package category

import "context"

type CategoryService interface {
	CreateCategory(context.Context, int) error
}
