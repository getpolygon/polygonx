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
package auth

import (
	"database/sql"
	"net/http"

	"github.com/getpolygon/corexp/internal/gen/postgres_codegen"
	"github.com/getpolygon/corexp/internal/settings"
	"github.com/go-chi/jwtauth"
	"github.com/lestrrat-go/jwx/jwt"
)

var verifier *jwtauth.JWTAuth

// This function will create a JWT authentication verifier which will
// be used by a middleware, provided by the `jwtauth` package.
func CreateAuthVerifier(s *settings.Settings) func() *jwtauth.JWTAuth {
	return func() *jwtauth.JWTAuth {
		// https://community.auth0.com/t/jwt-signing-algorithms-rs256-vs-hs256/7720
		verifier = jwtauth.New("HS256", []byte(s.Security.JWTSigningKey), nil)
		return verifier
	}
}

func Authenticator(postgres *postgres_codegen.Queries, s *settings.Settings) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// The source of the following code was provided by github.com/go-chi/jwtauth
			// more info at: https://github.com/go-chi/jwtauth/blob/v1.2.0/jwtauth.go#L163
			// Copyright (c) 2015-Present https://github.com/go-chi authors

			// MIT License

			// Permission is hereby granted, free of charge, to any person obtaining a copy of
			// this software and associated documentation files (the "Software"), to deal in
			// the Software without restriction, including without limitation the rights to
			// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
			// the Software, and to permit persons to whom the Software is furnished to do so,
			// subject to the following conditions:

			// The above copyright notice and this permission notice shall be included in all
			// copies or substantial portions of the Software.

			// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
			// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
			// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
			// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
			// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
			// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
			// --------------CODE SNIPPET START--------------
			token, _, err := jwtauth.FromContext(r.Context())
			if err != nil {
				http.Error(w, err.Error(), 401)
				return
			}

			if token == nil || jwt.Validate(token) != nil {
				http.Error(w, http.StatusText(401), 401)
				return
			}
			// --------------CODE SNIPPET END-----------------
			// Getting the ID of the user for further validation
			sub, ok := token.Get("sub")
			if !ok || sub == nil {
				http.Error(w, http.StatusText(401), 401)
				return
			}

			// Checking whether the user exists
			exists, err := postgres.CheckUserExistsByID(r.Context(), sub.(string))
			if err != nil {
				// If there are no rows returned by PostgreSQL -> User does not exist -> 401.
				if err == sql.ErrNoRows {
					http.Error(w, http.StatusText(401), 401)
					return
				}

				http.Error(w, http.StatusText(500), 500)
				return
			}

			// If the user with the ID provided in the JWT does not exist
			if !exists {
				http.Error(w, http.StatusText(401), 401)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// This function will attempt to retrieve user's ID from
// the JWT.
func GetUserIDFromRequest(r *http.Request) (string, error) {
	_, token, err := jwtauth.FromContext(r.Context())
	return token["sub"].(string), err
}

// This function will attempt to generate a JWT for a user by
// using their ID as the subject field.
func GenTokenWithUserID(id string) (string, error) {
	_, tok, err := verifier.Encode(map[string]interface{}{
		"sub": id,
	})

	return tok, err
}
