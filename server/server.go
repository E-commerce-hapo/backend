package server

import (
	"net/http"

	"github.com/E-commerce-hapo/backend/registry"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Server ...
type Server struct {
	categoryServer *CategoryServer
}

func New(r *registry.Registry) *Server {
	return &Server{
		categoryServer: newCategoryServer(r),
	}
}
func (s *Server) router() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	//apmgorilla.Instrument(r)
	r.Route("/api", func(r chi.Router) {
		r.Mount("/category", s.categoryServer.router())
	})

	return r
}
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router().ServeHTTP(w, r)
}
