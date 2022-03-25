package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	auth "polygon.am/core/api/v1/middleware/auth"
	"polygon.am/core/api/v1/routers"
)

// This router includes all routes that are being used by the v1
// API of Polygon.
func Router() *chi.Mux {
	r := chi.NewRouter()

	// Mounting authentication routes before the the jwt authorization
	// middleware to enable access to the auth routes without a token.
	r.Mount("/auth", routers.AuthRouter())

	r.Group(func(r chi.Router) {
		// Using JWT authentication on all API routes
		r.Use(jwtauth.Verifier(auth.Strategy))
		r.Use(jwtauth.Authenticator)

		r.Mount("/users", routers.UsersRouter())
		r.Mount("/posts", routers.PostsRouter())
	})

	return r
}
