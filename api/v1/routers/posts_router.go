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
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"polygon.am/core/pkg/persistence"
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

func GetPostById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, "invalid id")
		return
	}

	post, err := persistence.Queries.GetPostByID(context.Background(), int32(id))
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, "internal error: "+err.Error())
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, post)
	return
}

func DeletePostById(w http.ResponseWriter, r *http.Request) {}

func UpdatePostById(w http.ResponseWriter, r *http.Request) {}

func GetPostsOfUserByUsername(w http.ResponseWriter, r *http.Request) {}
