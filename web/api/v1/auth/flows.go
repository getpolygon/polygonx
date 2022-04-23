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
	"errors"
	"net/http"

	"gitea.com/go-chi/binding"
	"github.com/getpolygon/corexp/internal/deps"
	"github.com/getpolygon/corexp/internal/gen/postgres_codegen"
	"github.com/go-chi/render"
	"golang.org/x/crypto/bcrypt"
)

// This route is used for confirming user registration. It is only
// enabled if `polygon.security.accounts.forceEmailVerification` is
// forced from the configuration, and the required options in the
// config have valid values.
func SignUpConfirmation() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

// The sign in route is where the users will be able to and
// consume the API via the provided jwt token send in the
// JSON response.
func SignIn(deps *deps.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body := new(SignInRequestBody)
		// Parsing the values from the request payload to the newly created
		// body variable, with corresponding values. This function, however,
		// does not validate the request payload.
		if err := binding.Bind(r, body); err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err)
			return
		}

		user, err := deps.Postgres.GetFullUserByEmail(r.Context(), body.Email)
		if err != nil {
			// If nothing was returned by the SQL query, that means
			// that the user does not exist, however, the error will
			// need to be handled, since we are not getting the user
			// as a `nil`.
			if errors.Is(err, sql.ErrNoRows) {
				render.Status(r, http.StatusForbidden)
				render.JSON(w, r, "user not found")
				return
			}

			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		passwordHash, providedPass := []byte(user.Password), []byte(body.Password)
		// Comparing the passwords provided by the user and the
		// hashed one, stored in the database.
		err = bcrypt.CompareHashAndPassword(passwordHash, providedPass)
		if err != nil {
			// Comparison fails and the passwords do not match.
			render.Status(r, http.StatusForbidden)
			render.JSON(w, r, "invalid password.")
			return
		}

		// Generating a JWT for the user, using only their ID
		// for the "sub" field.
		token, err := GenTokenWithUserID(user.ID)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, token)
	}
}

// This route is used for creating an account for the users. It
// executes instantly, without any need for email validation if
// the `polygon.security.accounts.forceEmailVerification` option
// in the config is either unspecified or is of value `false`.
func SignUp(d *deps.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, body := r.Context(), new(SignUpRequestBody)
		if err := binding.Bind(r, body); err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err)
			return
		}

		pld := postgres_codegen.GetUserByEmailOrUsernameParams{
			Email:    body.Email,
			Username: body.Username,
		}
		if _, err := d.Postgres.GetUserByEmailOrUsername(ctx, pld); err != nil && !errors.Is(err, sql.ErrNoRows) {
			render.Status(r, http.StatusForbidden)
			render.JSON(w, r, "user already exists: "+err.Error())
			return
		} else {
			// TODO: Not implemented.
		}
	}
}
