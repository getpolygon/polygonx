package v1

import (
	"github.com/go-chi/chi/v5"
	"polygon.am/core/api/v1/routers"
)

// This router includes all routes that are being used by the v1
// API of Polygon.
func Router() *chi.Mux {
	router := chi.NewRouter()
	router.Mount("/users", routers.UsersRouter())

	return router
}
