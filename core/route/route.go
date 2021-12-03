package route

import (
	"net/http"

	log2 "github.com/kiem-toan/pkg/log"
	zipkinhttp "github.com/openzipkin/zipkin-go/middleware/http"
	"go.elastic.co/apm/module/apmgorilla"

	"github.com/kiem-toan/pkg/doc/swagger"

	"github.com/kiem-toan/domain/enums/api_type"

	"github.com/kiem-toan/pkg/auth"

	"github.com/gorilla/mux"
	"github.com/kiem-toan/cmd/audit-server/build"
	intzipkin "github.com/kiem-toan/pkg/zipkin"
)

type Route struct {
	Path        string
	HandlerFunc http.HandlerFunc
	Method      string
	Type        api_type.ApiType
}

func NewRouter(routes []Route) *mux.Router {
	router := mux.NewRouter()

	apmgorilla.Instrument(router)

	// Config Zipkin Tracing with Zipkin Middleware
	tracer, err := intzipkin.NewTracer()
	if err != nil {
		log2.Fatal(err, nil, nil)
	}
	router.Use(zipkinhttp.NewServerMiddleware(
		tracer,
		zipkinhttp.SpanName("Request")),
	)

	// Create Sub Router by API Type
	swaggerRouter := router.PathPrefix("/docs").Subrouter()
	internalAPIRouter := router.PathPrefix("/api").Subrouter()
	publicAPIRouter := router.PathPrefix("/").Subrouter()

	for _, route := range routes {
		handleFnc := auth.CORS(route.HandlerFunc)
		switch route.Type {
		case api_type.Swagger:
			swaggerRouter.Handle(route.Path, handleFnc).Methods(route.Method)
		case api_type.Internal:
			handleFnc = auth.TokenAuthMiddleware(route.HandlerFunc)
			internalAPIRouter.Handle(route.Path, handleFnc).Methods(route.Method)
		case api_type.Public:
			publicAPIRouter.Handle(route.Path, handleFnc).Methods(route.Method)
		}
	}
	return router
}

func AllRoutes(app *build.App) []Route {
	internalRoutes := []Route{
		// SWAGGER
		{"/", swagger.RedocHandler(), http.MethodGet, api_type.Swagger},
		{"/swagger.json", swagger.SwaggerHandler("/swagger.json"), http.MethodGet, api_type.Swagger},
		// CATEGORY
		{"/CreateCategory", app.CategoryHandler.CreateCategoryHandler, http.MethodPost, api_type.Internal},
		{"/ListCategories", app.CategoryHandler.ListCategoriesHandler, http.MethodPost, api_type.Internal},
		{"/CreateToken", app.CategoryHandler.CreateTokenHandler, http.MethodPost, api_type.Public},
		{"/VerifyToken", app.CategoryHandler.VerifyTokenHandler, http.MethodPost, api_type.Internal},
		{"/GetTokenData", app.CategoryHandler.GetTokenDataHandler, http.MethodPost, api_type.Internal},
		{"/RefreshToken", app.CategoryHandler.GetTokenDataHandler, http.MethodPost, api_type.Internal},
	}
	return internalRoutes
}
