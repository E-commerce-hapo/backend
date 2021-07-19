package category

type ICategoryService interface {
	CreateCategory(test *CreateCategoryRequest) (*CreateCategoryResponse, error)
}

// swagger:parameters Category CreateCategoryRequest
type CreateCategoryRequest struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Test Test `json:"test"`
}

// swagger:model Test
type Test struct {
	TestID string `json:"test"`
}

// swagger:response CreateCategoryResponse
type CreateCategoryResponse struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Obc CreateCategoryRequest `json:"obc"`
}