package category

import "context"

type CategoryRepositoryService interface {
	InsertCategory(context.Context, int) error
}
