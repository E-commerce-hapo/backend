package handler

import (
	"net/http"

	"github.com/E-commerce-hapo/backend/application/category"
	"github.com/E-commerce-hapo/backend/application/category/aggregate"
	"github.com/E-commerce-hapo/backend/application/category/query"
	"github.com/E-commerce-hapo/backend/pkg/auth"
	"github.com/E-commerce-hapo/backend/pkg/errorx"
	"github.com/E-commerce-hapo/backend/pkg/httpx"
	"github.com/E-commerce-hapo/backend/pkg/paging"
)

type CategoryHandler struct {
	CategoryAggr  aggregate.ICategoryAggr
	CategoryQuery query.ICategoryQuery
}

func New(categoryAggr aggregate.ICategoryAggr, categoryQuery query.ICategoryQuery) *CategoryHandler {
	return &CategoryHandler{
		CategoryAggr:  categoryAggr,
		CategoryQuery: categoryQuery,
	}
}

// @Tags 		 Category
// @Summary      Create Category
// @Description  Tạo danh mục (Category) từ thông tin client gửi lên
// @Accept       json
// @Produce      json
// @Param        CreateCategory  body      category2.CreateCategoryRequest  true  "Request CreateCategory"
// @Success      200          {object}  httpx.CreatedResponse
// @Router       /api/CreateCategory [post]
func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var t CreateCategoryRequest
	if err := httpx.ParseRequest(r, &t); err != nil {
		httpx.WriteError(ctx, w, err)
		return
	}
	err := h.CategoryAggr.CreateCategory(ctx, &category.CreateCategoryArgs{})
	if err != nil {
		httpx.WriteError(ctx, w, err)
		return
	}
	//httpx.WriteReponse(ctx, w, http.StatusOK, inter)
}

func (h *CategoryHandler) ListCategories(w http.ResponseWriter, r *http.Request) {
	var q ListCategoriesRequest
	if err := httpx.ParseRequest(r, &q); err != nil {
		httpx.WriteError(r.Context(), w, errorx.ErrInvalidParameter(err))
		return
	}
	response, err := h.CategoryQuery.ListCategories(r.Context(), &category.ListCategoriesArgs{Paging: &paging.Paging{
		Limit:  q.Paging.Limit,
		Offset: q.Paging.Offset,
		Sorts:  q.Paging.Sorts,
	}})
	if err != nil {
		httpx.WriteError(r.Context(), w, err)
		return
	}
	httpx.WriteReponse(r.Context(), w, http.StatusOK, response)
}

func (h *CategoryHandler) CreateTokenHandler(w http.ResponseWriter, r *http.Request) {
	var t CreateCategoryRequest
	if err := httpx.ParseRequest(r, &t); err != nil {
		httpx.WriteError(r.Context(), w, errorx.ErrInvalidParameter(err))
		return
	}
	token, err := auth.GenerateToken(1234)
	if err != nil {
		httpx.WriteError(r.Context(), w, err)
		return
	}
	httpx.WriteReponse(r.Context(), w, http.StatusOK, token)
}

func (h *CategoryHandler) VerifyTokenHandler(w http.ResponseWriter, r *http.Request) {
	var t CreateCategoryRequest
	if err := httpx.ParseRequest(r, &t); err != nil {
		httpx.WriteError(r.Context(), w, errorx.ErrInvalidParameter(err))
		return
	}
	claims, err := auth.GetCustomClaimsFromRequest(r)
	if err != nil {
		httpx.WriteError(r.Context(), w, err)
		return
	}
	httpx.WriteReponse(r.Context(), w, http.StatusOK, claims)
}

func (h *CategoryHandler) GetTokenDataHandler(w http.ResponseWriter, r *http.Request) {
	var t CreateCategoryRequest
	if err := httpx.ParseRequest(r, &t); err != nil {
		httpx.WriteError(r.Context(), w, errorx.ErrInvalidParameter(err))
		return
	}
	claims, err := auth.GetCustomClaimsFromRequest(r)
	if err != nil {
		httpx.WriteError(r.Context(), w, err)
		return
	}
	httpx.WriteReponse(r.Context(), w, http.StatusOK, claims)
}
