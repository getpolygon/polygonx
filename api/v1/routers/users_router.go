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
