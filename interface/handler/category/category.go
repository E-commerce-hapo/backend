package category

import (
	"net/http"

	category2 "github.com/kiem-toan/domain/api/category"

	"github.com/kiem-toan/infrastructure/httpx"
	"github.com/kiem-toan/interface/controller/category"
)

type CategoryHandler struct {
	CategoryService *category.CategoryService
}

func New(categorySvc *category.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		CategoryService: categorySvc,
	}
}

// swagger:route POST /"url" "danh mục" "kiểu request"
// swagger:route POST /CreateCategory  Category CreateCategoryRequest
// responses:
// 200: CreateCategoryResponse
func (h *CategoryHandler) CreateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var t *category2.CreateCategoryRequest
	if err := httpx.ParseRequest(r, &t); err != nil {
		httpx.RespondError(w, err)
	}
	inter, err := h.CategoryService.CreateCategory(nil, t)
	if err != nil {
		httpx.RespondError(w, err)
		return
	}
	httpx.RespondJSON(w, http.StatusOK, inter)
}

func (h *CategoryHandler) ListCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	var t *category2.CreateCategoryRequest
	if err := httpx.ParseRequest(r, &t); err != nil {
		httpx.RespondError(w, err)
	}
	inter, err := h.CategoryService.ListCategories(nil, t)
	if err != nil {
		httpx.RespondError(w, err)
		return
	}
	httpx.RespondJSON(w, http.StatusOK, inter)
}
