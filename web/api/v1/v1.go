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
package v1

import (
	"github.com/getpolygon/corexp/internal/deps"
	"github.com/getpolygon/corexp/web/api/v1/auth"
	"github.com/getpolygon/corexp/web/api/v1/posts"
	"github.com/getpolygon/corexp/web/api/v1/users"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
)

func Router(deps *deps.Dependencies) *chi.Mux {
	r := chi.NewRouter()

	// Mounting authentication routes before the the jwt authorization
	// middleware to enable access to the auth routes without a token.
	r.Mount("/auth", auth.Router(deps))

	// The rest of the routes will be protected and will need a JWT token
	// for access.
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(auth.CreateAuthVerifier(deps.Settings)))
		r.Use(auth.Authenticator(deps))

		r.Mount("/users", users.Router())
		r.Mount("/posts", posts.Router())
	})

	return r
}
