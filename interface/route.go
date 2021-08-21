package _interface

import (
	"net/http"

	"github.com/kiem-toan/infrastructure/auth"

	"github.com/gin-gonic/gin"

	"github.com/kiem-toan/cmd/audit-server/build"
)

type Route struct {
	Path        string
	HandlerFunc gin.HandlerFunc
	Method      string
}

func NewRouter(routes []Route) *gin.Engine {
	router := gin.New()
	// TODO: Swagger, finish later

	groupRouter := router.Group("/api")
	groupRouter.Use(auth.TokenAuthMiddleware())

	for _, route := range routes {
		handleFnc := route.HandlerFunc
		switch route.Method {
		case http.MethodGet:
			groupRouter.GET(route.Path, handleFnc)
		case http.MethodPost:
			groupRouter.POST(route.Path, handleFnc)
		case http.MethodPut:
			groupRouter.PUT(route.Path, handleFnc)
		case http.MethodDelete:
			groupRouter.DELETE(route.Path, handleFnc)
		}
	}
	return router
}

func AllRoutes(app *build.App) []Route {
	routes := []Route{
		// CATEGORY
		{"/CreateCategory", app.CategoryHandler.CreateCategoryHandler, http.MethodPost},
		{"/ListCategories", app.CategoryHandler.ListCategoriesHandler, http.MethodPost},
		//{"/CreateToken", app.CategoryHandler.CreateTokenHandler, http.MethodPost, nil},
		//{"/VerifyToken", app.CategoryHandler.VerifyTokenHandler, http.MethodPost, nil},
		//{"/GetTokenData", app.CategoryHandler.GetTokenDataHandler, http.MethodPost, nil},
		//{"/RefreshToken", app.CategoryHandler.GetTokenDataHandler, http.MethodPost, nil},
	}
	return routes
}
