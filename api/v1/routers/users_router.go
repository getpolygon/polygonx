package routers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// `UsersRouter()` exposes the routers related to users to
// the main router of a specific API version.
func UsersRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Put("/modify", UpdateUserById)
	r.Get("/with-id/{id}", GetUserById)
	r.Get("/{username}", GetUserByUsername)
	r.Delete("/close-account", CloseUserAccount)

	return r
}

// This endpoint is dedicated to fetching users' public
// information by using their usernames.
func GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	render.JSON(w, r, username)
}

// This endpoint is dedicated to fetching users' public
// information by using their IDs. This is useful when
// there are API consumers that would need to identify
// the user even when their username changes.
func GetUserById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	render.JSON(w, r, id)
}

// This endpoint is dedicated to closing users' accounts
// and deleting all the information that is related to
// their profiles.
func CloseUserAccount(w http.ResponseWriter, r *http.Request) {

}

// This endpoint is dedicated to updating users' public
// and private information by using their IDs.
func UpdateUserById(w http.ResponseWriter, r *http.Request) {

}
