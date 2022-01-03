package category

import (
	"time"
)

type CreatedCategoryEvent struct {
	Time time.Time
	ID   string
}
