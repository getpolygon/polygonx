package routers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// `AuthRouter()` exposes the routes related to authentication flow to the
// main router of a specific API version.
func AuthRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/signin", SignIn)
	r.Post("/signup", SignUp)

	return r
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, "sign in")
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, "sign up")
}
