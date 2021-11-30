package _interface

import (
	"net/http"

	"github.com/kiem-toan/domain/enums/api_type"

	"github.com/kiem-toan/pkg/auth"

	"github.com/gorilla/mux"
	"github.com/kiem-toan/cmd/audit-server/build"
)

type Route struct {
	Path        string
	HandlerFunc http.HandlerFunc
	Method      string
	Type        api_type.ApiType
}

func NewRouter(routes []Route) *mux.Router {
	router := mux.NewRouter()
	// TODO: Swagger, finish later
	//opts := middleware.RedocOpts{SpecURL: "/swagger.json"}
	//sh := middleware.Redoc(opts, nil)
	//router.Handle("/docs", sh).Methods(http.MethodGet)
	//router.Handle("/swagger.json", http.FileServer(http.Dir("./")))

	router = router.PathPrefix("/api").Subrouter()
	router.Use()
	for _, route := range routes {
		handleFnc := route.HandlerFunc
		if route.Type == api_type.Internal {
			handleFnc = auth.TokenAuthMiddleware(route.HandlerFunc)
		}
		router.Handle(route.Path, handleFnc).Methods(route.Method)
	}
	return router
}

func AllRoutes(app *build.App) []Route {
	routes := []Route{
		// CATEGORY
		{"/CreateCategory", app.CategoryHandler.CreateCategoryHandler, http.MethodPost, api_type.Internal},
		{"/ListCategories", app.CategoryHandler.ListCategoriesHandler, http.MethodPost, api_type.Internal},
		{"/CreateToken", app.CategoryHandler.CreateTokenHandler, http.MethodPost, api_type.Public},
		{"/VerifyToken", app.CategoryHandler.VerifyTokenHandler, http.MethodPost, api_type.Internal},
		{"/GetTokenData", app.CategoryHandler.GetTokenDataHandler, http.MethodPost, api_type.Internal},
		{"/RefreshToken", app.CategoryHandler.GetTokenDataHandler, http.MethodPost, api_type.Internal},
	}
	return routes
}
