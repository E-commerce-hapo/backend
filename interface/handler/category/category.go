package category

import (
	"net/http"

	category2 "github.com/kiem-toan/domain/api/category"

	"github.com/kiem-toan/interface/controller/category"
	"github.com/kiem-toan/pkg/httpx"
)

type CategoryHandler struct {
	CategoryService *category.CategoryService
}

func New(categorySvc *category.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		CategoryService: categorySvc,
	}
}

func (h *CategoryHandler) CreateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var t *category2.CreateCategoryRequest
	if err := httpx.ParseRequest(r, &t); err != nil {
		httpx.WriteError(ctx, w, err)
		return
	}
	inter, err := h.CategoryService.CreateCategory(ctx, t)
	if err != nil {
		httpx.WriteError(ctx, w, err)
		return
	}
	httpx.WriteReponse(ctx, w, http.StatusOK, inter)
}
