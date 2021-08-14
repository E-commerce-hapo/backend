package category

import (
	"net/http"

	"github.com/k0kubun/pp"

	"github.com/kiem-toan/infrastructure/idx"

	"github.com/kiem-toan/infrastructure/auth"

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
		pp.Println(err)
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
	pp.Println("c: ", c)
	response, err := h.CategoryService.ListCategories(c, &t)
	if err != nil {
		g.ResponseError(err)
		return
	}
	g.Response(http.StatusOK, response)
}

func (h *CategoryHandler) CreateTokenHandler(c *gin.Context) {
	g := httpx.Gin{C: c}
	var t category2.CreateCategoryRequest
	if err := g.ParseRequest(&t); err != nil {
		g.ResponseError(errorx.New(http.StatusInternalServerError, err, "Can not parse request"))
	}
	token, err := auth.GenerateToken(idx.NewID())
	if err != nil {
		g.ResponseError(err)
		return
	}
	g.Response(http.StatusOK, token)
}

func (h *CategoryHandler) VerifyTokenHandler(c *gin.Context) {
	g := httpx.Gin{C: c}
	var t category2.CreateCategoryRequest
	if err := g.ParseRequest(&t); err != nil {
		g.ResponseError(errorx.New(http.StatusInternalServerError, err, "Can not parse request"))
	}
	claims, err := auth.ExtractTokenMetaData(g.C.Request)
	if err != nil {
		g.ResponseError(err)
		return
	}
	g.Response(http.StatusOK, claims)
}

func (h *CategoryHandler) GetTokenDataHandler(c *gin.Context) {
	g := httpx.Gin{C: c}
	var t category2.CreateCategoryRequest
	if err := g.ParseRequest(&t); err != nil {
		g.ResponseError(errorx.New(http.StatusInternalServerError, err, "Can not parse request"))
	}
	claims, err := auth.ExtractTokenMetaData(g.C.Request)
	if err != nil {
		g.ResponseError(err)
		return
	}
	g.Response(http.StatusOK, claims)
}
