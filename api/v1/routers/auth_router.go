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
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/jackc/pgerrcode"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"polygon.am/core/api/v1/middleware/auth"
	"polygon.am/core/api/v1/middleware/validation"
	"polygon.am/core/pkg/persistence"
	"polygon.am/core/pkg/persistence/codegen"
)

// This struct contains the standard request body with its validations
// for the sign up endpoint.
type SignUpRequestBody struct {
	// An email address can contain at most 254 characters
	// more info: https://stackoverflow.com/questions/386294/what-is-the-maximum-length-of-a-valid-email-address#:~:text=%22There%20is%20a%20length%20limit,total%20length%20of%20320%20characters.
	Email    string `json:"email" validate:"required,email,max=254"`
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=8,alphanum"`
}

// This struct contains the standard request body with its validations
// for the sign in endpoint.
type SignInRequestBody struct {
	// An email address can contain at most 254 characters
	// more info: https://stackoverflow.com/questions/386294/what-is-the-maximum-length-of-a-valid-email-address#:~:text=%22There%20is%20a%20length%20limit,total%20length%20of%20320%20characters.
	Email    string `json:"email" validate:"required,email,max=254"`
	Password string `json:"password" validate:"required,min=8"`
}

func AuthRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/signin", SignIn)
	r.Post("/signup", SignUp)
	r.Post("/close-account", CloseAccount)
	r.Post("/signup/confirm", SignUpConfirmation)

	return r
}

// This route is used for confirming user registration. It is only
// enabled if `polygon.security.accounts.forceEmailVerification` is
// forced from the configuration, and the required options in the
// config have valid values.
func SignUpConfirmation(w http.ResponseWriter, r *http.Request) {

}

// The sign in route is where the users will be able to and
// consume the API via the provided jwt token send in the
// JSON response.
func SignIn(w http.ResponseWriter, r *http.Request) {
	body := new(SignInRequestBody)
	if err := validation.ValidateRequest(r, body); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err)
		return
	}

	user, err := persistence.Queries.GetFullUserByEmail(r.Context(), body.Email)
	if err != nil {
		if errors.As(err, &pq.Error{}) {
			switch err := err.(*pq.Error); err.Code {
			case pgerrcode.NoData, pgerrcode.NoDataFound:
				{
					render.Status(r, http.StatusNotFound)
					render.JSON(w, r, "user not found")
					return
				}
			default:
				{
					render.Status(r, http.StatusInternalServerError)
					render.JSON(w, r, err)
					return
				}
			}
		} else if err == sql.ErrNoRows {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, "user not found")
			return
		}

		render.Status(r, http.StatusNotImplemented)
		render.JSON(w, r, "unknown error")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		render.Status(r, http.StatusForbidden)
		render.JSON(w, r, "invalid password")
		return
	}

	token, err := auth.GenTokenWithUserID(user.ID)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, token)
	return
}

// This route is used for creating an account for the users. It
// executes instantly, without any need for email validation if
// the `polygon.security.accounts.forceEmailVerification` option
// in the config is either unspecified or is of value `false`.
func SignUp(w http.ResponseWriter, r *http.Request) {
	body := new(SignUpRequestBody)
	// Validating request form fields and sending an error if
	// the validation fails.
	if err := validation.ValidateRequest(r, body); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err)
		return
	}

	// Creating password hash from the original password and storing only
	// the encrypted one in the database.
	password, err := bcrypt.GenerateFromPassword([]byte(r.Form.Get("password")), bcrypt.DefaultCost)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, "password hashing error")
		return
	}

	// Persisting user's records in the database
	user, err := persistence.Queries.InsertUser(r.Context(), codegen.InsertUserParams{
		Name:     body.Name,
		Email:    body.Email,
		Username: body.Username,
		Password: string(password),
	})

	// Converting the error to pq Error and validating the code from
	// the PostgreSQL query via a helper library.
	if err != nil {
		switch err := err.(*pq.Error); err.Code {
		case pgerrcode.UniqueViolation:
			{
				render.Status(r, http.StatusForbidden)
				render.JSON(w, r, "user "+user.Username+" already exists")
				return
			}
		default:
			{
				render.Status(r, http.StatusInternalServerError)
				render.JSON(w, r, "unknown error")
				return
			}
		}
	}

	token, err := auth.GenTokenWithUserID(user.ID)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, err)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, token)
	return
}

// This route is used for deleting user accounts. The process is
// straightforward and does not need to be verified by email. It
// executes instantly, and deletes all the information associated
// with the user, including user settings, posts, comments, etc.
func CloseAccount(w http.ResponseWriter, r *http.Request) {

}
