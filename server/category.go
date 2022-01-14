package server

import (
	"github.com/go-chi/chi"

	categoryhandler "github.com/E-commerce-hapo/backend/application/category/handler"
	"github.com/E-commerce-hapo/backend/pkg/auth"
	"github.com/E-commerce-hapo/backend/registry"
)

type CategoryServer struct {
	categoryHandler *categoryhandler.CategoryHandler
}

func newCategoryServer(r *registry.Registry) *CategoryServer {
	return &CategoryServer{
		categoryHandler: categoryhandler.New(r.RegisterCategoryAggr(), r.RegisterCategoryQuery()),
	}
}

// Path: /api/Category/
func (s *CategoryServer) router() chi.Router {
	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Use(auth.TokenAuthMiddleware)
		r.Post("/CreateCategory", s.categoryHandler.CreateCategory)
	})

	r.Group(func(r chi.Router) {
		r.Get("/ListCategories", s.categoryHandler.ListCategories)
	})
	return r
}
