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
	"github.com/spf13/viper"
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
	// username := r.Form.Get("username")
	// password := r.Form.Get("password")

	// TODO: Find the user by their username or email address.
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	// name := r.Form.Get("name")
	// email := r.Form.Get("email")
	// username := r.Form.Get("username")
	// password := r.Form.Get("password")

	if viper.GetBool("polygon.security.accounts.forceEmailVerification") {
		// TODO: Handle user sign up, with email verification.
	}

	// TODO: Handle user sign up without email verification.
}
