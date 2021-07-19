package _interface

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kiem-toan/cmd/audit-server/build"
)

type Route struct {
	Path        string
	HandlerFunc http.HandlerFunc
	Method      string
}

func NewRouter(routes []Route) *mux.Router {
	router := mux.NewRouter()
	// TODO: Swagger, finish later
	//opts := middleware.RedocOpts{SpecURL: "/swagger.json"}
	//sh := middleware.Redoc(opts, nil)
	//router.Handle("/docs", sh).Methods(http.MethodGet)
	//router.Handle("/swagger.json", http.FileServer(http.Dir("./")))

	router = router.PathPrefix("/api").Subrouter()
	for _, route := range routes {
		handleFnc := route.HandlerFunc
		router.Handle(route.Path, handleFnc).Methods(route.Method)
	}
	return router
}

func AllRoutes(app *build.App) []Route {
	routes := []Route{
		{"/CreateCategory", app.CategoryHandler.CreateCategoryHandler, http.MethodPost},
	}
	return routes
}
