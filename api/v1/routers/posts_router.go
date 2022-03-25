package routers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// `PostsRouter()` exposes the routes related to posts to the
// main router of a specific API version.
func PostsRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/{id}", GetPostById)
	r.Delete("/{id}", DeletePostById)
	r.Put("/modify/{id}", UpdatePostById)
	r.Get("/of-user/{username}", GetPostsOfUserByUsername)

	return r
}

func GetPostById(w http.ResponseWriter, r *http.Request) {}

func DeletePostById(w http.ResponseWriter, r *http.Request) {}

func UpdatePostById(w http.ResponseWriter, r *http.Request) {}

func GetPostsOfUserByUsername(w http.ResponseWriter, r *http.Request) {}
