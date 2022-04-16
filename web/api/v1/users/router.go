package users

import "github.com/go-chi/chi/v5"

func Router() *chi.Mux {
	r := chi.NewRouter()
	r.Put("/modify", UpdateUserByID)
	r.Get("/with-id/{id}", GetUserByID)
	r.Get("/{username}", GetUserByUsername)

	return r
}
