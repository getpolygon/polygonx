// BSD 3-Clause License

// Copyright (c) 2021, Michael Grigoryan
// All rights reserved.

// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:

// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.

// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.

// 3. Neither the name of the copyright holder nor the names of its
//    contributors may be used to endorse or promote products derived from
//    this software without specific prior written permission.

// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
package routers

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"polygon.am/core/api/v1/middleware/auth"
	"polygon.am/core/api/v1/middleware/validation"
	"polygon.am/core/pkg/persistence"
	"polygon.am/core/pkg/persistence/codegen"
)

func PostsRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/{id}", GetPostByID)
	r.Post("/create", CreatePost)
	r.Delete("/{id}", DeletePostByID)
	r.Patch("/modify/{id}", ModifyPostByID)
	r.Get("/of/{username}", GetPostsOfUserByUsername)

	return r
}

type CreatePostRequestBody struct {
	Content string `json:"content"`
	Title   string `json:"title" validate:"max=80"`
}

// This endpoint is used for creating posts.
func CreatePost(w http.ResponseWriter, r *http.Request) {
	body := new(CreatePostRequestBody)
	if err := validation.ValidateRequest(r, body); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, "bad payload")
		return
	}

	user, err := auth.GetUserIDFromRequest(r)
	if err != nil {
		render.Status(r, http.StatusForbidden)
		render.JSON(w, r, "jwt error")
		return
	}

	post, err := persistence.Queries.InsertPost(r.Context(), codegen.InsertPostParams{
		User:    user,
		Title:   body.Title,
		Content: sql.NullString{String: body.Content, Valid: body.Content != ""},
	})

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, post)
	return
}

// This route is used for returning posts created by a certain
// user by specifying their username.
func GetPostsOfUserByUsername(w http.ResponseWriter, r *http.Request) {
	username, cursor := chi.URLParam(r, "username"), string(r.URL.Query().Get("cursor"))
	posts, err := persistence.Queries.GetPostsOfUserAfterID(r.Context(), codegen.GetPostsOfUserAfterIDParams{
		ID:       cursor,
		Username: username,
	})

	if err != nil {
		render.Status(r, http.StatusNotImplemented)
		render.JSON(w, r, "unknown error")
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, posts)
	return
}

// This route is used for fetching post information with the
// ID provided from URL parameters.
func GetPostByID(w http.ResponseWriter, r *http.Request) {

}

// This route is used for modifying post information with the
// ID provided from URL parameters.
func ModifyPostByID(w http.ResponseWriter, r *http.Request) {

}

// This route is used for deleting all information related to
// a specific post found by its ID. This will remove the post
// information itself, the comments, the upvotes and everything
// else related to it.
func DeletePostByID(w http.ResponseWriter, r *http.Request) {

}
