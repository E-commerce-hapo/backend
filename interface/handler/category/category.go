package category

import (
	"net/http"

	"github.com/kiem-toan/infrastructure/errorx"

	"github.com/gin-gonic/gin"

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

func (h *CategoryHandler) CreateCategoryHandler(c *gin.Context) {
	g := httpx.Gin{C: c}
	var t category2.CreateCategoryRequest
	if err := g.ParseRequest(&t); err != nil {
		g.ResponseError(errorx.New(http.StatusInternalServerError, err, "Can not parse request"))
		return
	}
	response, err := h.CategoryService.CreateCategory(nil, &t)
	if err != nil {
		g.ResponseError(err)
		return
	}
	g.Response(http.StatusOK, response)
}

func (h *CategoryHandler) ListCategoriesHandler(c *gin.Context) {
	g := httpx.Gin{C: c}
	var t category2.CreateCategoryRequest
	if err := g.ParseRequest(&t); err != nil {
		g.ResponseError(errorx.New(http.StatusInternalServerError, err, "Can not parse request"))
	}
	response, err := h.CategoryService.ListCategories(nil, &t)
	if err != nil {
		g.ResponseError(err)
		return
	}
	g.Response(http.StatusOK, response)
}
