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

//func NewRouter(routes []Route) *mux.Router {
//	router := mux.NewRouter()
//	// TODO: Swagger, finish later
//	//opts := middleware.RedocOpts{SpecURL: "/swagger.json"}
//	//sh := middleware.Redoc(opts, nil)
//	//router.Handle("/docs", sh).Methods(http.MethodGet)
//	//router.Handle("/swagger.json", http.FileServer(http.Dir("./")))
//
//	router = router.PathPrefix("/api").Subrouter()
//	for _, route := range routes {
//		handleFnc := route.HandlerFunc
//		router.Handle(route.Path, handleFnc).Methods(route.Method)
//	}
//	return router
//}

func NewRouter(routes []Route) *gin.Engine {
	router := gin.New()
	// TODO: Swagger, finish later
	//opts := middleware.RedocOpts{SpecURL: "/swagger.json"}
	//sh := middleware.Redoc(opts, nil)
	//router.Handle("/docs", sh).Methods(http.MethodGet)
	//router.Handle("/swagger.json", http.FileServer(http.Dir("./")))

	//router = router.PathPrefix("/api").Subrouter()
	//for _, route := range routes {
	//	handleFnc := route.HandlerFunc
	//	router.Handle(route.Path, handleFnc).Methods(route.Method)
	//}

	router2 := router.Group("/")
	router2.POST("/CreateToken", routes[2].HandlerFunc)
	router2.POST("/VerifyToken", routes[3].HandlerFunc)
	router2.POST("/GetTokenData", routes[4].HandlerFunc)
	router2.POST("/RefreshToken", routes[5].HandlerFunc)
	groupRouter := router.Group("/api")
	groupRouter.Use(auth.TokenAuthMiddleware())
	for ind, route := range routes {
		if ind <= 1 {
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

	}
	return router
}

func AllRoutes(app *build.App) []Route {
	routes := []Route{
		// CATEGORY
		{"/CreateCategory", app.CategoryHandler.CreateCategoryHandler, http.MethodPost},
		{"/ListCategories", app.CategoryHandler.ListCategoriesHandler, http.MethodPost},
		{"/CreateToken", app.CategoryHandler.CreateTokenHandler, http.MethodPost},
		{"/VerifyToken", app.CategoryHandler.VerifyTokenHandler, http.MethodPost},
		{"/GetTokenData", app.CategoryHandler.GetTokenDataHandler, http.MethodPost},
		{"/RefreshToken", app.CategoryHandler.GetTokenDataHandler, http.MethodPost},
	}
	return routes
}
