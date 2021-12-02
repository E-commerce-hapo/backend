package category

import (
	"net/http"

	category2 "github.com/kiem-toan/domain/api/category"
	"github.com/kiem-toan/pkg/auth"
	"github.com/kiem-toan/pkg/errorx"

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

// @Tags 		 Category
// @Summary      Create Category
// @Description  Tạo danh mục (Category) từ thông tin client gửi lên
// @Accept       json
// @Produce      json
// @Param        CreateCategory  body      category2.CreateCategoryRequest  true  "Request CreateCategory"
// @Success      200          {object}  httpx.CreatedResponse
// @Router       /api/CreateCategory [post]
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

func (h *CategoryHandler) ListCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	var t category2.ListCategoriesRequest
	if err := httpx.ParseRequest(r, &t); err != nil {
		httpx.WriteError(r.Context(), w, errorx.Errorf(http.StatusBadRequest, err, "Can not parse request"))
		return
	}
	response, err := h.CategoryService.ListCategories(r.Context(), &t)
	if err != nil {
		httpx.WriteError(r.Context(), w, err)
		return
	}
	httpx.WriteReponse(r.Context(), w, http.StatusOK, response)
}

func (h *CategoryHandler) CreateTokenHandler(w http.ResponseWriter, r *http.Request) {
	var t category2.CreateCategoryRequest
	if err := httpx.ParseRequest(r, &t); err != nil {
		httpx.WriteError(r.Context(), w, errorx.Errorf(http.StatusBadRequest, err, "Can not parse request"))
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
	var t category2.CreateCategoryRequest
	if err := httpx.ParseRequest(r, &t); err != nil {
		httpx.WriteError(r.Context(), w, errorx.Errorf(http.StatusBadRequest, err, "Can not parse request"))
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
	var t category2.CreateCategoryRequest
	if err := httpx.ParseRequest(r, &t); err != nil {
		httpx.WriteError(r.Context(), w, errorx.Errorf(http.StatusBadRequest, err, "Can not parse request"))
		return
	}
	claims, err := auth.GetCustomClaimsFromRequest(r)
	if err != nil {
		httpx.WriteError(r.Context(), w, err)
		return
	}
	httpx.WriteReponse(r.Context(), w, http.StatusOK, claims)
}
