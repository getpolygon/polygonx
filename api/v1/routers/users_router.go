package routers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// `UsersRouter()` exposes handlers, middleware and routers to the main
// router of a specific API version.
func UsersRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Put("/modify", UpdateUserById)
	router.Get("/with-id/{id}", GetUserById)
	router.Get("/{username}", GetUserByUsername)
	router.Delete("/close-account", CloseUserAccount)

	return router
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
